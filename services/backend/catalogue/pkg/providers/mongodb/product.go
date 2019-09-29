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
	productRepository struct {
		session *mgo.Session
	}
)

const (
	productCollection = "product"
)

var (
	config cfg.Config
)

func newProductRepository(session *mgo.Session) engine.ProductRepository {
	return &productRepository{session}
}

func (r productRepository) Insert(c context.Context, p domain.Product) error {
	s := r.session.Clone()
	defer s.Close()

	var product domain.Product

	product.Creator = p.Creator
	product.Description = p.Description
	product.Details = p.Details
	product.Managers = p.Managers
	product.Date = p.Date
	product.Followers = p.Followers
	product.Tags = p.Tags
	product.Type = p.Type
	product.Topic = p.Topic
	product.Date = p.Date

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Insert(&product)
	if err != nil {
		log.Println("insert error ", err)
	}

	return nil
}

func (r productRepository) Update(c context.Context, p domain.Product, id string) error {
	s := r.session.Clone()
	defer s.Close()

	var product domain.Product

	product.Creator = p.Creator
	product.Description = p.Description
	product.Details = p.Details
	product.Managers = p.Managers
	product.Date = p.Date
	product.Followers = p.Followers
	product.Tags = p.Tags
	product.Type = p.Type
	product.Topic = p.Topic
	product.Date = p.Date

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Update(bson.M{"_id": toBSON(id)}, p)
	if err != nil {
		return err
	}

	return nil
}

func (r productRepository) Query(c context.Context, query *engine.Query) []*domain.Product {
	s := r.session.Clone()
	defer s.Close()

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	p := []*domain.Product{}
	q := translateQuery(col, query)
	q.All(&p)

	return nil
}

func (r productRepository) FindOne(c context.Context, id string) (*domain.Product, error) {
	s := r.session.Clone()
	defer s.Close()

	var product *domain.Product
	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"_id": toBSON(id)}).One(&product)

	if err != nil {
		return nil, nil
	}

	return product, nil
}

// ListAllProducts used for finding all user products
// by the passed skip and take parameters
func (r productRepository) ListAllProducts(ctx context.Context, skip uint64, take uint64) ([]domain.Product, error) {
	s := r.session.Clone()
	defer s.Close()
	log.Println("got in repo")

	var products []domain.Product

	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(nil).All(&products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r productRepository) Remove(c context.Context, id string) (string, error) {
	s := r.session.Clone()
	defer s.Close()

	err := s.DB(config.MongoDB).C(config.MongoCollection).Remove(bson.M{"_id": toBSON(id)}) //.One(&product)

	if err != nil {
		return "", nil
	}

	return "", nil
}
