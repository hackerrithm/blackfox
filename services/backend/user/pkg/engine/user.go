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

	"github.com/hackerrithm/blackfox/services/backend/user/pkg/domain"
)

type (
	// User ...
	User interface {
		// RegisterUser used for registering of new user accounts
		RegisterUser(ctx context.Context, user domain.Register) (*domain.User, error)
		// UpdateUser used for update of a specific user account
		UpdateUser(ctx context.Context, user domain.User) (*domain.User, error)
		// PutUser used for ... user account
		PutUser(ctx context.Context, user domain.User) (*domain.User, error)
		// GetUserByID used for finding of a specific user account
		// by the passed identifier
		GetUserByID(ctx context.Context, id string) (*domain.User, error)
		// GetUserByUserName used for finding of a specific user account
		// by the passed username
		GetUserByUserName(ctx context.Context, username string) (*domain.User, error)
		// GetUserByEmailAddress used for finding of a specific user account
		// by the passed email address
		GetUserByEmailAddress(ctx context.Context, email string) (*domain.User, error)
		// LoginUser used for sign in of a user by authenticating
		// credentials username and password
		LoginUser(ctx context.Context, loginRequest domain.Login) (*domain.User, error)
		// ListAllUsers used for finding all user accounts
		// by the passed skip and take parameters
		ListAllUsers(ctx context.Context, skip uint64, take uint64) ([]domain.User, error)
		// DeleteUserAccount used for removing/deleting of a specific user account
		// by the passed user identifier
		DeleteUserAccount(ctx context.Context, id string) (bool, error)
		//GenerateToken ...
		GenerateToken(ctx context.Context, id string) (string, error)
		// ParseToken ...
		ParseToken(ctx context.Context, token string) (map[string]interface{}, error)
	}

	user struct {
		repository UserRepository
		jwt        JWTSignParser
	}
)

var (
	userInstance User
	userOnce     sync.Once
)

func (f *engineFactory) NewUser() User {
	userOnce.Do(func() {
		userInstance = &user{
			repository: f.NewUserRepository(),
			jwt:        f.NewAuthentication(),
		}
	})
	return userInstance
}
