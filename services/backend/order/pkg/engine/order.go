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
	"sync"

	"golang.org/x/net/context"

	"github.com/hackerrithm/blackfox/services/backend/order/pkg/domain"
)

type (
	// Order ...
	Order interface {
		Insert(ctx context.Context, userID, details, description, typ string, totalPrice float64, products []domain.OrderedProduct) (*domain.Order, error)

		// Update is the update-a-order use-case
		Update(ctx context.Context, id, userID, details, description, typ string, totalPrice float64, products []domain.OrderedProduct) error

		// FindOne ...
		FindOne(ctx context.Context, id string) (*domain.Order, error)

		// RemoveDelete ...
		Remove(ctx context.Context, id string) (string, error)

		// ListAllOrders ...
		ListAllOrders(ctx context.Context, skip, take uint64) (*[]domain.Order, error)
	}

	order struct {
		repository OrderRepository
	}
)

var (
	orderInstance Order
	orderOnce     sync.Once
)

func (f *engineFactory) NewOrder() Order {
	orderOnce.Do(func() {
		orderInstance = &order{
			repository: f.NewOrderRepository(),
			// jwt:        f.JWTSignParser,
		}
	})
	return orderInstance
}
