package authentication

import (
	"context"
	"log"
	"strings"

	"net/http"

	"github.com/hackerrithm/blackfox/services/backend/api/pkg"
	apiCfg "github.com/hackerrithm/blackfox/services/backend/api/pkg/configs"
	auth "github.com/hackerrithm/blackfox/services/backend/auth/cmd/auth/client"
	user "github.com/hackerrithm/blackfox/services/backend/user/cmd/user/client"
	"github.com/kelseyhightower/envconfig"
)

var contxt context.Context
var secretkey = "12This98Is34A76String56Used65As78Secret01"
var userObj pkg.User

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type authServer struct {
	authClient *auth.Client
	userClient *user.Client
}

// getToken gets Authorization key from headers
func getToken(request *http.Request, response http.ResponseWriter, next http.Handler) (string, error) {
	log.Println(" -------------------- >>>>>>>>>>>>>> IN getToken method")
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		log.Println("header is empty")
		next.ServeHTTP(response, request)
	}

	log.Println(" -------------------- >>>>>>>>>>>>>> IN getToken method :::::: header not empty")

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		log.Println("invalid auth token")
	}
	log.Println(" -------------------- >>>>>>>>>>>>>> IN getToken method with value ", authHeaderParts[1])

	return authHeaderParts[1], nil
}

// AuthHandlerMiddleware ...
func AuthHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		var cfg apiCfg.Config
		err := envconfig.Process("", &cfg)
		if err != nil {
			log.Fatal("------------>>>>>>>>>>>>>>>> error setting up config", err)
		}

		// return authHeaderParts[1], nil
		header := request.Header.Get("Authorization")

		switch header {
		case "":
			next.ServeHTTP(response, request)
		default:
			authHeader := request.Header.Get("Authorization")
			if authHeader == "" {
				log.Println("header is empty")
				next.ServeHTTP(response, request)
			}

			authHeaderParts := strings.Split(authHeader, " ")
			if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
				log.Println("invalid auth token")
			}

			tokenStr := authHeaderParts[1] //, err := getToken(request, response, next)
			// if err != nil {
			// 	log.Println("error on getting token")
			// }

			if tokenStr == "" {
				next.ServeHTTP(response, request)
			}

			// authClient, err := auth.NewClient(cfg.AuthServiceURL)
			// if err != nil {
			// 	log.Println("error connecting to user client")
			// 	authClient.Close()
			// }

			// userID, err := authClient.GetUserFromToken(contxt, tokenStr)
			// if err != nil {
			// 	log.Println("error parsing token")
			// }
			// authClient.Close()

			userClient, err := user.NewClient(cfg.UserServiceURL)
			if err != nil {
				log.Println("error connecting to user client")
				userClient.Close()
			}

			userID := "5dade4266700b1a5e37555ef"

			// get the user from the database
			user, err := userClient.GetUser(contxt, userID)
			if err != nil {
				log.Println("error getting user with token and all")
			}

			ctxt := context.WithValue(request.Context(), userCtxKey, user)
			next.ServeHTTP(response, request.WithContext(ctxt))

		}

		// if header == "" {
		// } else {

		// }
	})
}

// UserIDFromToken ...
// func UserIDFromToken(tokenString string) int {
// 	token, err := JwtDecode(tokenString)
// 	if err != nil {
// 		return 0
// 	}
// 	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
// 		if claims == nil {
// 			return 0
// 		}
// 		return claims.UserId
// 	} else {
// 		return 0
// 	}
// }

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *pkg.User {
	// log.Println("pkg user in forContext: ", &user)
	// raw, _ := ctx.Value(userCtxKey).(*pkg.User)
	// return raw
	raw := ctx.Value(userCtxKey)
	log.Println("rawwwwwww ---------->>>>>> ", raw)
	if raw == nil {
		return nil
	}
	return raw.(*pkg.User)
}
