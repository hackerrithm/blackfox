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

package engine

import (
	"context"
	"log"
	"time"

	"github.com/hackerrithm/blackfox/services/backend/auth/pkg/domain"
)

const (
	secretKey = "12This98Is34A76String56Used65As78Secret01"
)

// RegisterUser used for registering of new user users
func (s *auth) RegisterUser(ctx context.Context, user domain.Register) (*domain.User, error) {
	usr := domain.NewUser("", user.UserName, user.Password, user.FirstName, user.LastName, user.EmailAddress)
	return usr, nil
}

// LoginUser used for sign in of a user by authenticating
// credentials username and password
func (s *auth) LoginUser(ctx context.Context, user domain.Login) (*domain.User, error) {
	return nil, nil
}

func (s *auth) GenerateToken(ctx context.Context, id string) (string, error) {
	log.Println("generated token (ID)-- ", id)
	claims := map[string]interface{}{
		"userid": id,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	log.Println("generated token (claims)-- ", claims)
	return s.jwt.Sign(claims, secretKey)
}

func (s *auth) ParseToken(ctx context.Context, token string) (map[string]interface{}, error) {
	return s.jwt.Parse(token, secretKey)
}
