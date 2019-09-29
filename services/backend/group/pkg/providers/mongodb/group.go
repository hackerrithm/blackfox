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

	cfg "github.com/hackerrithm/blackfox/services/backend/group/configs"
	"github.com/hackerrithm/blackfox/services/backend/group/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/group/pkg/engine"
)

type (
	groupRepository struct {
		session *mgo.Session
	}
)

const (
	groupCollection = "group"
)

var (
	config cfg.Config
)

func newGroupRepository(session *mgo.Session) engine.GroupRepository {
	return &groupRepository{session}
}

func (r groupRepository) Insert(c context.Context, p domain.Group) error {
	s := r.session.Clone()
	defer s.Close()

	var group domain.Group

	group.Description = p.Description
	group.Details = p.Details
	group.Date = p.Date
	group.Title = p.Title
	group.People = p.People
	group.Type = p.Type

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Insert(&group)
	if err != nil {
		log.Println("insert error ", err)
	}

	return nil
}

func (r groupRepository) Update(c context.Context, p domain.Group, id string) error {
	s := r.session.Clone()
	defer s.Close()

	var group domain.Group

	group.Description = p.Description
	group.Details = p.Details
	group.Date = p.Date
	group.Title = p.Title
	group.People = p.People
	group.Type = p.Type

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Update(bson.M{"_id": toBSON(id)}, p)
	if err != nil {
		return err
	}

	return nil
}

func (r groupRepository) Query(c context.Context, query *engine.Query) []*domain.Group {
	s := r.session.Clone()
	defer s.Close()

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	p := []*domain.Group{}
	q := translateQuery(col, query)
	q.All(&p)

	return nil
}

func (r groupRepository) FindOne(c context.Context, id string) (*domain.Group, error) {
	s := r.session.Clone()
	defer s.Close()

	var group *domain.Group
	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"_id": toBSON(id)}).One(&group)

	if err != nil {
		return nil, nil
	}

	return group, nil
}

// ListAllGroups used for finding all user groups
// by the passed skip and take parameters
func (r groupRepository) ListAllGroups(ctx context.Context, skip uint64, take uint64) ([]domain.Group, error) {
	s := r.session.Clone()
	defer s.Close()
	log.Println("got in repo")

	var groups []domain.Group

	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(nil).All(&groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (r groupRepository) Remove(c context.Context, id string) (string, error) {
	s := r.session.Clone()
	defer s.Close()

	err := s.DB(config.MongoDB).C(config.MongoCollection).Remove(bson.M{"_id": toBSON(id)}) //.One(&group)

	if err != nil {
		return "", nil
	}

	return "", nil
}
