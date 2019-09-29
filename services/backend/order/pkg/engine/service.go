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

	"gopkg.in/mgo.v2/bson"

	"github.com/hackerrithm/blackfox/services/backend/order/pkg/domain"
)

func (s *order) Insert(ctx context.Context, userID, details, description, typ string, totalPrice float64, products []domain.OrderedProduct) (*domain.Order, error) {
	o := &domain.Order{
		UserID:   bson.ObjectIdHex(userID),
		Products: products,
	}
	// Calculate total price
	o.TotalPrice = 0.0
	for _, p := range products {
		o.TotalPrice += p.Price * float64(p.Quantity)
	}

	a := domain.NewOrder(userID, details, description, typ, o.TotalPrice, products)
	if err := s.repository.Insert(ctx, *a, products); err != nil {
		return nil, err
	}

	return a, nil
}

func (s *order) Update(ctx context.Context, id, userID, details, description, typ string, totalPrice float64, products []domain.OrderedProduct) error {
	return nil
}

func (s *order) FindOne(ctx context.Context, id string) (*domain.Order, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *order) ListAllOrders(ctx context.Context, skip, take uint64) (*[]domain.Order, error) {
	res, err := s.repository.ListAllOrders(ctx, skip, take)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *order) Remove(ctx context.Context, id string) (string, error) {
	return s.repository.Remove(ctx, id)
}
