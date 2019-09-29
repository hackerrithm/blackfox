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

	cfg "github.com/hackerrithm/blackfox/services/backend/match/configs"
	"github.com/hackerrithm/blackfox/services/backend/match/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/match/pkg/engine"
)

type (
	matchRepository struct {
		session *mgo.Session
	}
)

const (
	matchCollection = "match"
)

var (
	config cfg.Config
)

func newMatchRepository(session *mgo.Session) engine.MatchRepository {
	return &matchRepository{session}
}

func (r matchRepository) Insert(c context.Context, p domain.Match) error {
	s := r.session.Clone()
	defer s.Close()

	var match domain.Match

	match.Description = p.Description
	match.Details = p.Details
	match.Date = p.Date
	match.PersonID = p.PersonID
	match.Similarities = p.Similarities
	match.Type = p.Type

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Insert(&match)
	if err != nil {
		log.Println("insert error ", err)
	}

	return nil
}

func (r matchRepository) Update(c context.Context, p domain.Match, id string) error {
	s := r.session.Clone()
	defer s.Close()

	var match domain.Match

	match.Description = p.Description
	match.Details = p.Details
	match.Date = p.Date
	match.PersonID = p.PersonID
	match.Similarities = p.Similarities
	match.Type = p.Type

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Update(bson.M{"_id": toBSON(id)}, p)
	if err != nil {
		return err
	}

	return nil
}

func (r matchRepository) Query(c context.Context, query *engine.Query) []*domain.Match {
	s := r.session.Clone()
	defer s.Close()

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	p := []*domain.Match{}
	q := translateQuery(col, query)
	q.All(&p)

	return nil
}

func (r matchRepository) FindOne(c context.Context, id string) (*domain.Match, error) {
	s := r.session.Clone()
	defer s.Close()

	var match *domain.Match
	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"_id": toBSON(id)}).One(&match)

	if err != nil {
		return nil, nil
	}

	return match, nil
}

// ListAllMatches used for finding all user matchs
// by the passed skip and take parameters
func (r matchRepository) ListAllMatches(ctx context.Context, skip uint64, take uint64) ([]domain.Match, error) {
	s := r.session.Clone()
	defer s.Close()
	log.Println("got in repo")

	var matchs []domain.Match

	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(nil).All(&matchs)
	if err != nil {
		return nil, err
	}

	return matchs, nil
}

func (r matchRepository) Remove(c context.Context, id string) (string, error) {
	s := r.session.Clone()
	defer s.Close()

	err := s.DB(config.MongoDB).C(config.MongoCollection).Remove(bson.M{"_id": toBSON(id)}) //.One(&match)

	if err != nil {
		return "", nil
	}

	return "", nil
}
