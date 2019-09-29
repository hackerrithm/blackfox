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

	"github.com/hackerrithm/blackfox/services/backend/catalogue/pkg/domain"
)

type (
	// Catalogue ...
	Catalogue interface {
		Insert(ctx context.Context, name, details, description, typ string, products, tags []string) error

		// Update is the update-a-catalogue use-case
		Update(ctx context.Context, id, name, details, description, typ string, products, tags []string) error

		// Query is the list-the-catalogues use-case
		Query(ctx context.Context, skip uint64, take uint64) ([]domain.Catalogue, error)

		// FindOne ...
		FindOne(ctx context.Context, id string) (*domain.Catalogue, error)

		// RemoveDelete ...
		Remove(ctx context.Context, id string) (string, error)

		// ListAllCatalogues ...
		ListAllCatalogues(ctx context.Context, skip, take uint64) (*[]domain.Catalogue, error)
	}

	catalogue struct {
		repository CatalogueRepository
	}

	// Product ...
	Product interface {
		InsertProduct(ctx context.Context, name, details, description, typ string, tags []string, price float64, discount float32) error

		// Update is the update-a-catalogue use-case
		UpdateProduct(ctx context.Context, id, name, details, description, typ string, tags []string, price float64, discount float32) error

		// Query is the list-the-catalogues use-case
		QueryProduct(ctx context.Context, skip uint64, take uint64) ([]domain.Product, error)

		// FindOne ...
		FindOneProduct(ctx context.Context, id string) (*domain.Product, error)

		// RemoveDelete ...
		RemoveProduct(ctx context.Context, id string) (string, error)

		// ListAllProducts ...
		ListAllProducts(ctx context.Context, skip, take uint64) (*[]domain.Product, error)
	}

	product struct {
		repository ProductRepository
	}
)

var (
	catalogueInstance Catalogue
	productInstance   Product

	catalogueOnce sync.Once
	productOnce   sync.Once
)

func (f *engineFactory) NewCatalogue() Catalogue {
	catalogueOnce.Do(func() {
		catalogueInstance = &catalogue{
			repository: f.NewCatalogueRepository(),
		}
	})
	return catalogueInstance
}

func (f *engineFactory) NewProduct() Product {
	catalogueOnce.Do(func() {
		productInstance = &product{
			repository: f.NewProductRepository(),
		}
	})
	return productInstance
}
