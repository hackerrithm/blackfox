// Copyright 2019 kemar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	log1 "log"
	"os"
	"time"

	"cloud.google.com/go/profiler"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/tinrab/retry"
	"contrib.go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"

	goalCfg "github.com/hackerrithm/blackfox/services/backend/goal/configs"
	"github.com/hackerrithm/blackfox/services/backend/goal/pkg/adapters"
	"github.com/hackerrithm/blackfox/services/backend/goal/pkg/engine"
	"github.com/hackerrithm/blackfox/services/backend/goal/pkg/providers/mongodb"
)

type (
	goal struct {
		engine.Goal
	}
)

var log *logrus.Logger

func init() {
	log = logrus.New()
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
}

func main() {
	var cfg goalCfg.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var strg engine.StorageFactory
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		strg, err = mongodb.NewStorage(
			cfg.MongoHost,
			cfg.MongoDB,
			cfg.MongoUsername,
			cfg.MongoPassword,
			cfg.MongoCollection,
			cfg.MONGOURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer strg.Close()

	eObj := engine.NewEngine(strg)
	log1.Println(strg, eObj)
	goalObject := &goal{eObj.NewGoal()}
	log1.Println("Listening on port 8080...")
	log1.Fatal(adapters.ListenGRPC(goalObject, 8080))
}

func initJaegerTracing() {
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
			ServiceName: "goalservice",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(exporter)
	log.Info("jaeger initialization completed.")
}

func initStats(exporter *stackdriver.Exporter) {
	view.SetReportingPeriod(60 * time.Second)
	view.RegisterExporter(exporter)
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Warn("Error registering default server views")
	} else {
		log.Info("Registered default server views")
	}
}

func initStackDriverTracing() {
	// TODO(ahmetb) this method is duplicated in other microservices using Go
	// since they are not sharing packages.
	for i := 1; i <= 3; i++ {
		exporter, err := stackdriver.NewExporter(stackdriver.Options{})
		if err != nil {
			log.Warnf("failed to initialize stackdriver exporter: %+v", err)
		} else {
			trace.RegisterExporter(exporter)
			trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
			log.Info("registered stackdriver tracing")

			// Register the views to collect server stats.
			initStats(exporter)
			return
		}
		d := time.Second * 10 * time.Duration(i)
		log.Infof("sleeping %v to retry initializing stackdriver exporter", d)
		time.Sleep(d)
	}
	log.Warn("could not initialize stackdriver exporter after retrying, giving up")
}

func initTracing() {
	initJaegerTracing()
	initStackDriverTracing()
}

func initProfiling(service, version string) {
	// TODO: (ahmetb) this method is duplicated in other microservices using Go
	// since they are not sharing packages.
	for i := 1; i <= 3; i++ {
		if err := profiler.Start(profiler.Config{
			Service:        service,
			ServiceVersion: version,
			// ProjectID must be set if not running on GCP.
			// ProjectID: "my-project",
		}); err != nil {
			log.Warnf("failed to start profiler: %+v", err)
		} else {
			log.Info("started stackdriver profiler")
			return
		}
		d := time.Second * 10 * time.Duration(i)
		log.Infof("sleeping %v to retry initializing stackdriver profiler", d)
		time.Sleep(d)
	}
	log.Warn("could not initialize stackdriver profiler after retrying, giving up")
}
