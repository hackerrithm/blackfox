package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/profiler"
	"contrib.go.opencensus.io/exporter/jaeger"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"github.com/99designs/gqlgen/handler"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/cors"

	multiserviceMap "github.com/hackerrithm/blackfox/services/backend/api/pkg/resolver"
	apiCfg "github.com/hackerrithm/blackfox/services/backend/api/pkg/configs"
)

const defaultPort = "8080"

func initJaegerTracing(log logrus.FieldLogger) {

	svcAddr := os.Getenv("JAEGER_SERVICE_ADDR")
	if svcAddr == "" {
		log.Info("jaeger initialization disabled.")
		return
	}

	// Register the Jaeger exporter to be able to retrieve
	// the collected spans.
	exporter, err := jaeger.NewExporter(jaeger.Options{
		Endpoint: fmt.Sprintf("http://%s", svcAddr),
		Process: jaeger.Process{
			ServiceName: "frontend",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(exporter)
	log.Info("jaeger initialization completed.")
}

func initStats(log logrus.FieldLogger, exporter *stackdriver.Exporter) {
	view.SetReportingPeriod(60 * time.Second)
	view.RegisterExporter(exporter)
	if err := view.Register(ochttp.DefaultServerViews...); err != nil {
		log.Warn("Error registering http default server views")
	} else {
		log.Info("Registered http default server views")
	}
	if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
		log.Warn("Error registering grpc default client views")
	} else {
		log.Info("Registered grpc default client views")
	}
}

func initStackdriverTracing(log logrus.FieldLogger) {
	// TODO(ahmetb) this method is duplicated in other microservices using Go
	// since they are not sharing packages.
	for i := 1; i <= 3; i++ {
		log = log.WithField("retry", i)
		exporter, err := stackdriver.NewExporter(stackdriver.Options{})
		if err != nil {
			// log.Warnf is used since there are multiple backends (stackdriver & jaeger)
			// to store the traces. In production setup most likely you would use only one backend.
			// In that case you should use log.Fatalf.
			log.Warnf("failed to initialize stackdriver exporter: %+v", err)
		} else {
			trace.RegisterExporter(exporter)
			log.Info("registered stackdriver tracing")

			// Register the views to collect server stats.
			initStats(log, exporter)
			return
		}
		d := time.Second * 20 * time.Duration(i)
		log.Debugf("sleeping %v to retry initializing stackdriver exporter", d)
		time.Sleep(d)
	}
	log.Warn("could not initialize stackdriver exporter after retrying, giving up")
}

func initTracing(log logrus.FieldLogger) {
	// This is a demo app with low QPS. trace.AlwaysSample() is used here
	// to make sure traces are available for observation and analysis.
	// In a production environment or high QPS setup please use
	// trace.ProbabilitySampler set at the desired probability.
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	initJaegerTracing(log)
	initStackdriverTracing(log)

}

func initProfiling(log logrus.FieldLogger, service, version string) {
	// TODO: (ahmetb) this method is duplicated in other microservices using Go
	// since they are not sharing packages.
	for i := 1; i <= 3; i++ {
		log = log.WithField("retry", i)
		if err := profiler.Start(profiler.Config{
			Service:        service,
			ServiceVersion: version,
			// ProjectID must be set if not running on GCP.
			// ProjectID: "my-project",
		}); err != nil {
			log.Warnf("warn: failed to start profiler: %+v", err)
		} else {
			log.Info("started stackdriver profiler")
			return
		}
		d := time.Second * 10 * time.Duration(i)
		log.Debugf("sleeping %v to retry initializing stackdriver profiler", d)
		time.Sleep(d)
	}
	log.Warn("warning: could not initialize stackdriver profiler after retrying, giving up")
}

func main() {
	var cfg apiCfg.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339Nano,
	}
	log.Out = os.Stdout

	go initProfiling(log, "apigateway", "1.0.0")
	go initTracing(log)

	s, err := multiserviceMap.NewGraphQLServer(cfg.UserServiceURL, cfg.AuthServiceURL, cfg.PostServiceURL, cfg.SpaceServiceURL,
		cfg.TaskServiceURL, cfg.ProfileServiceURL, cfg.GeographyServiceURL,
		cfg.GoalServiceURL, cfg.MatchServiceURL, cfg.RedisServiceURL)
	if err != nil {
		log.Fatal(err)
	}

	// router := chi.NewRouter()
	// mux := http.NewServeMux()
	router := mux.NewRouter()
	// router := chi.NewRouter()
	// router.Use(authentication.AuthHandlerMiddleware)
	// h := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowCredentials: false,
	// 	Debug:            true,
	// })
	h := cors.AllowAll()
	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	// router.Handle("/query", http.Handler(authentication.AuthHandlerMiddleware(handler.GraphQL(s.ToExecutableSchema(),
	// 	handler.WebsocketUpgrader(websocket.Upgrader{
	// 		CheckOrigin: func(r *http.Request) bool {
	// 			return true
	// 		},
	// 	})))))
	router.Handle("/query", h.Handler(handler.GraphQL(s.ToExecutableSchema(),
		handler.WebsocketUpgrader(websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}))))

	c := h.Handler(router)

	log.Println("listening graphql server from port:", defaultPort, cfg.Server.Port)
	log.Fatal(http.ListenAndServe(":"+defaultPort, c))
}
