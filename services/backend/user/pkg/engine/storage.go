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

	"github.com/hackerrithm/blackfox/services/backend/user/pkg/domain"
)

type (
	// UserRepository acts as interface for user specific storage
	// functions for actions against database layer
	UserRepository interface {
		// InsertUser used for inserts of new user accounts
		InsertUser(ctx context.Context, user *domain.User) error
		// FindUserByUsernameAndPassword used to do user authentication
		FindUserByUsernameAndPassword(ctx context.Context, username, password string) (*domain.User, error)
		// UpdateUser used for update of a specific user account
		UpdateUser(ctx context.Context, user domain.User) error
		// PutUser used for ... user account
		PutUser(ctx context.Context, user domain.User) error
		// FindUserByID used for finding of a specific user account
		// by the passed identifier
		FindUserByID(ctx context.Context, id string) (*domain.User, error)
		// FindUserByUserName used for finding of a specific user account
		// by the passed username
		FindUserByUserName(ctx context.Context, username string) (*domain.User, error)
		// FindUserByEmailAddress used for finding of a specific user account
		// by the passed email address
		FindUserByEmailAddress(ctx context.Context, email string) (*domain.User, error)
		// ListAllUsers used for finding all user accounts
		// by the passed skip and take parameters
		ListAllUsers(ctx context.Context, skip uint64, take uint64) ([]domain.User, error)
		// RemoveUserByID used for removing/deleting of a specific user account
		// by the passed unique identifier
		RemoveUserByID(ctx context.Context, id string) (bool, error)
	}

	// StorageFactory acts as interface for storage methods
	// performing actions against our database
	StorageFactory interface {
		// NewUserRepository returns a storage specific
		// UserRepository implementation
		NewUserRepository() UserRepository

		Close()
	}
)
