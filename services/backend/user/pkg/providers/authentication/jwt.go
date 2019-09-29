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
	"log"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/hackerrithm/blackfox/services/backend/user/pkg/engine"
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
func (j *jwt) Sign(claims map[string]interface{}, secretKey string) (map[string]interface{}, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, jwtgo.MapClaims(claims))
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println("StatusUnauthorized ", err)
	}

	returnObjectMap["token"] = tokenString

	return returnObjectMap, nil
}

// Parse ..
func (j *jwt) Parse(tokenStr, secret string) (map[string]interface{}, error) {
	token, err := jwtgo.Parse(tokenStr, func(token *jwtgo.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		e, ok := err.(*jwtgo.ValidationError)
		if ok {
			return nil, e
		}
		return nil, err
	}

	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}
