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

package authentication

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/hackerrithm/blackfox/services/backend/auth/pkg/engine"
)

type jwt struct {
	// token jwtgo.Token
}

// NewJWT ...
func newJWT( /*token jwtgo.Token*/ ) engine.JWTSignParser {
	return &jwt{ /*token*/ }
}

var returnObjectMap = make(map[string]interface{})

// Sign ...
func (j *jwt) Sign(claims map[string]interface{}, secretKey string) (string, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, jwtgo.MapClaims(claims))
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println("StatusUnauthorized ", err)
	}

	// returnObjectMap["token"] = tokenString

	return tokenString, nil
}

// Parse ..
func (j *jwt) Parse(tokenStr, secret string) error {
	// token, err := jwtgo.Parse(tokenStr, func(token *jwtgo.Token) (interface{}, error) {
	// 	return []byte(secret), nil
	// })
	// if err != nil {
	// 	e, ok := err.(*jwtgo.ValidationError)
	// 	if ok {
	// 		return nil, e
	// 	}
	// 	return nil, err
	// }

	// claims, ok := token.Claims.(jwtgo.MapClaims)
	// if !ok {
	// 	return nil, err
	// }
	// return claims, nil
	token, err := jwtgo.Parse(tokenStr, func(token *jwtgo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwtgo.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

// ExtractTokenID ...
func (j *jwt) ExtractTokenID(tokenString string) (string, error) {

	token, err := jwtgo.Parse(tokenString, func(token *jwtgo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtgo.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwtgo.MapClaims)
	if ok && token.Valid {
		userID := fmt.Sprintf("%v", claims["user_id"])
		if err != nil {
			return "", err
		}
		return userID, nil
	}
	return "", nil
}

//Pretty display the claims licely in the terminal
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}
