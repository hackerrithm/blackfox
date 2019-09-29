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

	cfg "github.com/hackerrithm/blackfox/services/backend/order/configs"
	"github.com/hackerrithm/blackfox/services/backend/order/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/order/pkg/engine"
)

type (
	orderRepository struct {
		session *mgo.Session
	}
)

const (
	orderCollection = "order"
)

var (
	config cfg.Config
)

func newOrderRepository(session *mgo.Session) engine.OrderRepository {
	return &orderRepository{session}
}

func (r orderRepository) Insert(c context.Context, p domain.Order, q []domain.OrderedProduct) error {
	s := r.session.Clone()
	defer s.Close()

	var order domain.Order

	order.Description = p.Description
	order.Details = p.Details
	order.CreatedOn = p.CreatedOn
	order.Type = p.Type
	order.TotalPrice = p.TotalPrice
	order.UserID = p.UserID
	for _, i := range q {
		order.Products = append(order.Products, i)
	}

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Insert(&order)
	if err != nil {
		log.Println("insert error ", err)
	}

	return nil
}

func (r orderRepository) Update(c context.Context, p domain.Order, id string) error {
	return nil
}

func (r orderRepository) FindOne(c context.Context, id string) (*domain.Order, error) {
	s := r.session.Clone()
	defer s.Close()

	var order *domain.Order
	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"_id": toBSON(id)}).One(&order)

	if err != nil {
		return nil, nil
	}

	return order, nil
}

// ListAllOrders used for finding all user orders
// by the passed skip and take parameters
func (r orderRepository) ListAllOrders(ctx context.Context, skip uint64, take uint64) ([]domain.Order, error) {
	s := r.session.Clone()
	defer s.Close()
	log.Println("got in repo")

	var orders []domain.Order

	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(nil).All(&orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (r orderRepository) Remove(c context.Context, id string) (string, error) {
	s := r.session.Clone()
	defer s.Close()

	err := s.DB(config.MongoDB).C(config.MongoCollection).Remove(bson.M{"_id": toBSON(id)}) //.One(&order)

	if err != nil {
		return "", nil
	}

	return "", nil
}
