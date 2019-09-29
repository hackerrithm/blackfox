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

	cfg "github.com/hackerrithm/blackfox/services/backend/catalogue/configs"
	"github.com/hackerrithm/blackfox/services/backend/catalogue/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/catalogue/pkg/engine"
)

type (
	catalogueRepository struct {
		session *mgo.Session
	}
)

const (
	catalogueCollection = "catalogue"
)

var (
	config cfg.Config
)

func newCatalogueRepository(session *mgo.Session) engine.CatalogueRepository {
	return &catalogueRepository{session}
}

func (r catalogueRepository) Insert(c context.Context, p domain.Catalogue) error {
	s := r.session.Clone()
	defer s.Close()

	var catalogue domain.Catalogue

	catalogue.Creator = p.Creator
	catalogue.Description = p.Description
	catalogue.Details = p.Details
	catalogue.Managers = p.Managers
	catalogue.Date = p.Date
	catalogue.Followers = p.Followers
	catalogue.Tags = p.Tags
	catalogue.Type = p.Type
	catalogue.Topic = p.Topic
	catalogue.Date = p.Date

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Insert(&catalogue)
	if err != nil {
		log.Println("insert error ", err)
	}

	return nil
}

func (r catalogueRepository) Update(c context.Context, p domain.Catalogue, id string) error {
	s := r.session.Clone()
	defer s.Close()

	var catalogue domain.Catalogue

	catalogue.Creator = p.Creator
	catalogue.Description = p.Description
	catalogue.Details = p.Details
	catalogue.Managers = p.Managers
	catalogue.Date = p.Date
	catalogue.Followers = p.Followers
	catalogue.Tags = p.Tags
	catalogue.Type = p.Type
	catalogue.Topic = p.Topic
	catalogue.Date = p.Date

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Update(bson.M{"_id": toBSON(id)}, p)
	if err != nil {
		return err
	}

	return nil
}

func (r catalogueRepository) Query(c context.Context, query *engine.Query) []*domain.Catalogue {
	s := r.session.Clone()
	defer s.Close()

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	p := []*domain.Catalogue{}
	q := translateQuery(col, query)
	q.All(&p)

	return nil
}

func (r catalogueRepository) FindOne(c context.Context, id string) (*domain.Catalogue, error) {
	s := r.session.Clone()
	defer s.Close()

	var catalogue *domain.Catalogue
	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"_id": toBSON(id)}).One(&catalogue)

	if err != nil {
		return nil, nil
	}

	return catalogue, nil
}

// ListAllCatalogues used for finding all user catalogues
// by the passed skip and take parameters
func (r catalogueRepository) ListAllCatalogues(ctx context.Context, skip uint64, take uint64) ([]domain.Catalogue, error) {
	s := r.session.Clone()
	defer s.Close()
	log.Println("got in repo")

	var catalogues []domain.Catalogue

	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(nil).All(&catalogues)
	if err != nil {
		return nil, err
	}

	return catalogues, nil
}

func (r catalogueRepository) Remove(c context.Context, id string) (string, error) {
	s := r.session.Clone()
	defer s.Close()

	err := s.DB(config.MongoDB).C(config.MongoCollection).Remove(bson.M{"_id": toBSON(id)}) //.One(&catalogue)

	if err != nil {
		return "", nil
	}

	return "", nil
}
