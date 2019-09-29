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

package mongodb

import (
	"log"

	"golang.org/x/net/context"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	cfg "github.com/hackerrithm/blackfox/services/backend/user/configs"
	"github.com/hackerrithm/blackfox/services/backend/user/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/user/pkg/engine"
)

type (
	userRepository struct {
		session *mgo.Session
	}
)

const (
	userCollection = "user"
)

var (
	config cfg.Config
)

func newUserRepository(session *mgo.Session) engine.UserRepository {
	return &userRepository{session}
}

// InsertUser used for inserts of new user accounts
func (r userRepository) InsertUser(ctx context.Context, u *domain.User) error {
	s := r.session.Clone()
	defer s.Close()

	var user domain.User

	user.Username = u.Username
	user.SetPassword(u.Password)
	user.Firstname = u.Firstname
	user.Lastname = u.Lastname

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Insert(&user)
	if err != nil {
		return err
	}

	log.Println("inserted i user: ", &user)

	return nil
}

func (r userRepository) FindUserByUsernameAndPassword(ctx context.Context, username, password string) (*domain.User, error) {
	s := r.session.Clone()
	defer s.Close()

	var user *domain.User

	log.Println("repo: ", config.MongoDB, " ", config.MongoCollection)
	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return nil, err
	}

	if !user.IsCredentialsVerified(password, user.Password) {
		return nil, nil
	}

	return user, nil
}

// UpdateUser used for update of a specific user account
func (r userRepository) UpdateUser(ctx context.Context, user domain.User) error {
	return nil
}

// PutUser used for ... user account
func (r userRepository) PutUser(ctx context.Context, user domain.User) error {
	return nil
}

// FindUserByID used for finding of a specific user account
// by the passed identifier
func (r userRepository) FindUserByID(ctx context.Context, id string) (*domain.User, error) {
	s := r.session.Clone()
	defer s.Close()

	var user *domain.User
	bsonid := bson.ObjectIdHex(id)
	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"_id": bsonid}).One(&user)
	if err != nil {
		return &domain.User{}, nil
	}
	return user, nil

}

// FindUserByUserName used for finding of a specific user account
// by the passed username
func (r userRepository) FindUserByUserName(ctx context.Context, username string) (*domain.User, error) {
	s := r.session.Clone()
	defer s.Close()

	var user *domain.User

	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"username": username}).One(&user)

	if err != nil {
		return &domain.User{}, nil
	}

	return user, nil

}

// FindUserByEmailAddress used for finding of a specific user account
// by the passed email address
func (r userRepository) FindUserByEmailAddress(ctx context.Context, email string) (*domain.User, error) {
	s := r.session.Clone()
	defer s.Close()

	var user *domain.User

	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"email_address": email}).One(&user)

	if err != nil {
		return &domain.User{}, nil
	}

	return user, nil

}

// ListAllUsers used for finding all user accounts
// by the passed skip and take parameters
func (r userRepository) ListAllUsers(ctx context.Context, skip uint64, take uint64) ([]domain.User, error) {
	s := r.session.Clone()
	defer s.Close()

	var users []domain.User

	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(nil).All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// RemoveUserByID used for removing/deleting of a specific user account
// by the passed unique identifier
func (r userRepository) RemoveUserByID(ctx context.Context, id string) (bool, error) {
	s := r.session.Clone()
	defer s.Close()

	bsonid := bson.ObjectIdHex(id)
	err := s.DB(config.MongoDB).C(config.MongoCollection).Remove(bson.M{"_id": bsonid})
	if err != nil {
		return false, err
	}

	return true, nil
}
