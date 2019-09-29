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

	cfg "github.com/hackerrithm/blackfox/services/backend/profile/configs"
	"github.com/hackerrithm/blackfox/services/backend/profile/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/profile/pkg/engine"
)

type (
	profileRepository struct {
		session *mgo.Session
	}
)

const (
	profileCollection = "profile"
)

var (
	config cfg.Config
)

func newProfileRepository(session *mgo.Session) engine.ProfileRepository {
	return &profileRepository{session}
}

func (r profileRepository) Insert(c context.Context, p domain.Profile) error {
	s := r.session.Clone()
	defer s.Close()

	var profile domain.Profile

	profile.About = p.About
	profile.UserName = p.UserName
	profile.BackgroundImage = p.BackgroundImage
	profile.ProfileImage = p.ProfileImage
	profile.Followers = p.Followers
	profile.Level = p.Level
	profile.Rings = p.Rings
	profile.DateLastUpdated = p.DateLastUpdated

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Insert(&profile)
	if err != nil {
		log.Println("insert error ", err)
	}

	return nil
}

func (r profileRepository) Update(c context.Context, p domain.Profile, id string) error {
	s := r.session.Clone()
	defer s.Close()

	var profile domain.Profile

	profile.About = p.About
	profile.UserName = p.UserName
	profile.BackgroundImage = p.BackgroundImage
	profile.ProfileImage = p.ProfileImage
	profile.Followers = p.Followers
	profile.Level = p.Level
	profile.Rings = p.Rings
	profile.DateLastUpdated = p.DateLastUpdated

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Update(bson.M{"_id": toBSON(id)}, p)
	if err != nil {
		return err
	}

	return nil
}

func (r profileRepository) Query(c context.Context, query *engine.Query) []*domain.Profile {
	s := r.session.Clone()
	defer s.Close()

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	p := []*domain.Profile{}
	q := translateQuery(col, query)
	q.All(&p)

	return nil
}

func (r profileRepository) FindOne(c context.Context, id string) (*domain.Profile, error) {
	s := r.session.Clone()
	defer s.Close()

	var profile *domain.Profile
	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"_id": toBSON(id)}).One(&profile)

	if err != nil {
		return nil, nil
	}

	return profile, nil
}

// ListAllProfiles used for finding all user profiles
// by the passed skip and take parameters
func (r profileRepository) ListAllProfiles(ctx context.Context, skip uint64, take uint64) ([]domain.Profile, error) {
	s := r.session.Clone()
	defer s.Close()
	log.Println("got in repo")

	var profiles []domain.Profile

	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(nil).All(&profiles)
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

func (r profileRepository) Remove(c context.Context, id string) (string, error) {
	s := r.session.Clone()
	defer s.Close()

	err := s.DB(config.MongoDB).C(config.MongoCollection).Remove(bson.M{"_id": toBSON(id)}) //.One(&profile)

	if err != nil {
		return "", nil
	}

	return "", nil
}
