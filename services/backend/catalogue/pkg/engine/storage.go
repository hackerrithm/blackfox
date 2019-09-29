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
	"golang.org/x/net/context"

	"github.com/hackerrithm/blackfox/services/backend/catalogue/pkg/domain"
)

type (
	// CatalogueRepository defines the methods that any
	// data storage provider needs to implement to get
	// and store catalogues
	CatalogueRepository interface {

		// Catalogue adds a new Catalogue to the datastore
		Insert(c context.Context, catalogue domain.Catalogue) error

		// Put adds a new Catalogue to the datastore
		Update(c context.Context, catalogue domain.Catalogue, id string) error

		// Query returns existing catalogues matching the
		// query provided
		Query(c context.Context, query *Query) []*domain.Catalogue

		// FindOne returns ...
		FindOne(c context.Context, id string) (*domain.Catalogue, error)

		// Remove ...
		Remove(c context.Context, id string) (string, error)

		// ListAllCatalogues ...
		ListAllCatalogues(ctx context.Context, skip uint64, take uint64) ([]domain.Catalogue, error)
	}

	// ProductRepository defines the methods that any
	// data storage provider needs to implement to get
	// and store product
	ProductRepository interface {

		// Product adds a new Product to the datastore
		Insert(c context.Context, product domain.Product) error

		// Put adds a new Product to the datastore
		Update(c context.Context, product domain.Product, id string) error

		// Query returns existing products matching the
		// query provided
		Query(c context.Context, query *Query) []*domain.Product

		// FindOne returns ...
		FindOne(c context.Context, id string) (*domain.Product, error)

		// Remove ...
		Remove(c context.Context, id string) (string, error)

		// ListAllProducts ...
		ListAllProducts(ctx context.Context, skip uint64, take uint64) ([]domain.Product, error)
	}

	// StorageFactory is the interface that a storage
	// provider needs to implement so that the engine can
	// request repository instances as it needs them
	StorageFactory interface {
		// NewCatalogueRepository returns a storage specific
		// CatalogueRepository implementation
		NewCatalogueRepository() CatalogueRepository

		// NewProductRepository returns a storage specific
		// ProductRepository implementation
		NewProductRepository() ProductRepository

		Close()
	}
)
