package authentication

import (
	"context"
	"log"
	"time"

	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/hackerrithm/blackfox/services/backend/api/pkg"
)

var contxt context.Context
var secretkey = []byte("secret_key")
var user pkg.User

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// AuthHandlerMiddleware ...
func AuthHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		header := request.Header.Get("Authorization")
		if header == "" {
			next.ServeHTTP(response, request)
		} else {
			token := jwt.New(jwt.SigningMethodHS256)
			token.Claims = jwt.MapClaims{
				"userid":   user.ID,
				"username": user.Username,
				"exp":      time.Now().Add(time.Hour * 24).Unix(),
			}
			tokenstring, err := token.SignedString(secretkey)
			if err != nil {
				log.Fatal("Error while generating token ", err)
			}
			ctxt := context.WithValue(request.Context(), userCtxKey, tokenstring)
			next.ServeHTTP(response, request.WithContext(ctxt))
		}
	})
}
