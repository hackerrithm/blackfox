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

	cfg "github.com/hackerrithm/blackfox/services/backend/reaction/configs"
	"github.com/hackerrithm/blackfox/services/backend/reaction/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/reaction/pkg/engine"
)

type (
	reactionRepository struct {
		session *mgo.Session
	}
)

const (
	reactionCollection = "reaction"
)

var (
	config cfg.Config
)

func newReactionRepository(session *mgo.Session) engine.ReactionRepository {
	return &reactionRepository{session}
}

func (r reactionRepository) Insert(c context.Context, p domain.Reaction) error {
	s := r.session.Clone()
	defer s.Close()

	var reaction domain.Reaction

	reaction.Description = p.Description
	reaction.Details = p.Details
	reaction.Date = p.Date
	reaction.PersonID = p.PersonID
	reaction.Type = p.Type

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Insert(&reaction)
	if err != nil {
		log.Println("insert error ", err)
	}

	return nil
}

func (r reactionRepository) Update(c context.Context, p domain.Reaction, id string) error {
	s := r.session.Clone()
	defer s.Close()

	var reaction domain.Reaction

	reaction.Description = p.Description
	reaction.Details = p.Details
	reaction.Date = p.Date
	reaction.PersonID = p.PersonID
	reaction.Type = p.Type

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Update(bson.M{"_id": toBSON(id)}, p)
	if err != nil {
		return err
	}

	return nil
}

func (r reactionRepository) Query(c context.Context, query *engine.Query) []*domain.Reaction {
	s := r.session.Clone()
	defer s.Close()

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	p := []*domain.Reaction{}
	q := translateQuery(col, query)
	q.All(&p)

	return nil
}

func (r reactionRepository) FindOne(c context.Context, id string) (*domain.Reaction, error) {
	s := r.session.Clone()
	defer s.Close()

	var reaction *domain.Reaction
	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"_id": toBSON(id)}).One(&reaction)

	if err != nil {
		return nil, nil
	}

	return reaction, nil
}

// ListAllReactiones used for finding all user reactions
// by the passed skip and take parameters
func (r reactionRepository) ListAllReactiones(ctx context.Context, skip uint64, take uint64) ([]domain.Reaction, error) {
	s := r.session.Clone()
	defer s.Close()
	log.Println("got in repo")

	var reactions []domain.Reaction

	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(nil).All(&reactions)
	if err != nil {
		return nil, err
	}

	return reactions, nil
}

func (r reactionRepository) Remove(c context.Context, id string) (string, error) {
	s := r.session.Clone()
	defer s.Close()

	err := s.DB(config.MongoDB).C(config.MongoCollection).Remove(bson.M{"_id": toBSON(id)}) //.One(&reaction)

	if err != nil {
		return "", nil
	}

	return "", nil
}
