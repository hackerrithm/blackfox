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
	"sync"

	"golang.org/x/net/context"

	"github.com/hackerrithm/blackfox/services/backend/auth/pkg/domain"
)

type (
	// Auth ...
	Auth interface {
		// RegisterUser used for registering of new user accounts
		RegisterUser(ctx context.Context, user domain.Register) (*domain.User, error)
		// LoginUser used for sign in of a user by authenticating
		// credentials username and password
		LoginUser(ctx context.Context, loginRequest domain.Login) (*domain.User, error)
		//GenerateToken ...
		GenerateToken(ctx context.Context, id string) (string, error)
		// ParseToken ...
		ParseToken(ctx context.Context, token string) error
		// ExtractTokenID ...
		ExtractTokenID(tokenString string) (string, error)
	}

	auth struct {
		jwt JWTSignParser
	}
)

var (
	authInstance Auth
	authOnce     sync.Once
)

func (f *engineFactory) NewUser() Auth {
	authOnce.Do(func() {
		authInstance = &auth{
			jwt: f.NewAuthentication(),
		}
	})
	return authInstance
}
