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

	cfg "github.com/hackerrithm/blackfox/services/backend/space/configs"
	"github.com/hackerrithm/blackfox/services/backend/space/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/space/pkg/engine"
)

type (
	spaceRepository struct {
		session *mgo.Session
	}
)

const (
	spaceCollection = "space"
)

var (
	config cfg.Config
)

func newSpaceRepository(session *mgo.Session) engine.SpaceRepository {
	return &spaceRepository{session}
}

func (r spaceRepository) Insert(c context.Context, p domain.Space) error {
	s := r.session.Clone()
	defer s.Close()

	var space domain.Space

	space.Creator = p.Creator
	space.Description = p.Description
	space.Details = p.Details
	space.Managers = p.Managers
	space.Date = p.Date
	space.Followers = p.Followers
	space.Tags = p.Tags
	space.Type = p.Type
	space.Topic = p.Topic
	space.Date = p.Date

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Insert(&space)
	if err != nil {
		log.Println("insert error ", err)
	}

	return nil
}

func (r spaceRepository) Update(c context.Context, p domain.Space, id string) error {
	s := r.session.Clone()
	defer s.Close()

	var space domain.Space

	space.Creator = p.Creator
	space.Description = p.Description
	space.Details = p.Details
	space.Managers = p.Managers
	space.Date = p.Date
	space.Followers = p.Followers
	space.Tags = p.Tags
	space.Type = p.Type
	space.Topic = p.Topic
	space.Date = p.Date

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Update(bson.M{"_id": toBSON(id)}, p)
	if err != nil {
		return err
	}

	return nil
}

func (r spaceRepository) Query(c context.Context, query *engine.Query) []*domain.Space {
	s := r.session.Clone()
	defer s.Close()

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	p := []*domain.Space{}
	q := translateQuery(col, query)
	q.All(&p)

	return nil
}

func (r spaceRepository) FindOne(c context.Context, id string) (*domain.Space, error) {
	s := r.session.Clone()
	defer s.Close()

	var space *domain.Space
	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"_id": toBSON(id)}).One(&space)

	if err != nil {
		return nil, nil
	}

	return space, nil
}

// ListAllSpaces used for finding all user spaces
// by the passed skip and take parameters
func (r spaceRepository) ListAllSpaces(ctx context.Context, skip uint64, take uint64) ([]domain.Space, error) {
	s := r.session.Clone()
	defer s.Close()
	log.Println("got in repo")

	var spaces []domain.Space

	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(nil).All(&spaces)
	if err != nil {
		return nil, err
	}

	return spaces, nil
}

func (r spaceRepository) Remove(c context.Context, id string) (string, error) {
	s := r.session.Clone()
	defer s.Close()

	err := s.DB(config.MongoDB).C(config.MongoCollection).Remove(bson.M{"_id": toBSON(id)}) //.One(&space)

	if err != nil {
		return "", nil
	}

	return "", nil
}
