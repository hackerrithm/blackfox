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

	"github.com/hackerrithm/blackfox/services/backend/user/pkg/domain"
)

const (
	secretKey = "12This98Is34A76String56Used65As78Secret01"
)

// RegisterUser used for registering of new user users
func (s *user) RegisterUser(ctx context.Context, user domain.Register) (*domain.User, error) {
	usr := domain.NewUser("", user.UserName, user.Password, user.FirstName, user.LastName, user.EmailAddress)
	err := s.repository.InsertUser(ctx, usr)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

// LoginUser used for sign in of a user by authenticating
// credentials username and password
func (s *user) LoginUser(ctx context.Context, user domain.Login) (*domain.User, error) {
	usr, err := s.repository.FindUserByUsernameAndPassword(ctx, user.UserName, user.Password)
	if err != nil {
		return nil, err
	}
	return usr, nil

}

// UpdateUser used for update of a specific user user
func (s *user) UpdateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	return nil, nil
}

// PutUser used for ... user user
func (s *user) PutUser(ctx context.Context, user domain.User) (*domain.User, error) {
	return nil, nil
}

// GetUserByID used for finding of a specific user user
// by the passed identifier
func (s *user) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	u, err := s.repository.FindUserByID(ctx, id)
	if err != nil {
		return &domain.User{}, err
	}
	return u, nil
}

// GetUserByUserName used for finding of a specific user user
// by the passed username
func (s *user) GetUserByUserName(ctx context.Context, username string) (*domain.User, error) {
	u, err := s.repository.FindUserByUserName(ctx, username)
	if err != nil {
		return &domain.User{}, err
	}
	return u, nil
}

// GetUserByEmailAddress used for finding of a specific user user
// by the passed email address
func (s *user) GetUserByEmailAddress(ctx context.Context, email string) (*domain.User, error) {
	u, err := s.repository.FindUserByID(ctx, email)
	if err != nil {
		return &domain.User{}, err
	}
	return u, nil
}

// ListAllUsers used for finding all user users
// by the passed skip and take parameters
func (s *user) ListAllUsers(ctx context.Context, skip uint64, take uint64) ([]domain.User, error) {
	u, err := s.repository.ListAllUsers(ctx, 4, 8)
	if err != nil {
		return []domain.User{}, err
	}
	return u, nil
}

// DeleteUserAccount used for removing/deleting of a specific user user
// by the passed user identifier
func (s *user) DeleteUserAccount(ctx context.Context, id string) (bool, error) {
	b, err := s.repository.RemoveUserByID(ctx, id)
	if err != nil {
		return b, err
	}
	return b, nil
}

func (s *user) GenerateToken(ctx context.Context, id string) (string, error) {
	log.Println("generated token (ID)-- ", id)
	claims := map[string]interface{}{
		"userid": id,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	log.Println("generated token (claims)-- ", claims)
	return s.jwt.Sign(claims, secretKey)
}

func (s *user) ParseToken(ctx context.Context, token string) (map[string]interface{}, error) {
	return s.jwt.Parse(token, secretKey)
}
