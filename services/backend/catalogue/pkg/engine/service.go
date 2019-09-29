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

	"github.com/hackerrithm/blackfox/services/backend/catalogue/pkg/domain"
)

func (s *catalogue) Insert(ctx context.Context, name, details, description, typ string, products, tags []string) error {
	a := domain.NewCatalogue(name, details, description, typ, products, tags)
	if err := s.repository.Insert(ctx, *a); err != nil {
		return err
	}

	return nil
}

func (s *catalogue) Update(ctx context.Context, id, name, details, description, typ string, products, tags []string) error {
	a := domain.NewCatalogue(name, details, description, typ, products, tags)
	if err := s.repository.Update(ctx, *a, id); err != nil {
		return err
	}

	return nil
}

func (s *catalogue) FindOne(ctx context.Context, id string) (*domain.Catalogue, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *catalogue) ListAllCatalogues(ctx context.Context, skip, take uint64) (*[]domain.Catalogue, error) {
	res, err := s.repository.ListAllCatalogues(ctx, skip, take)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *catalogue) Query(ctx context.Context, skip uint64, take uint64) ([]domain.Catalogue, error) {
	return nil, nil
}

func (s *catalogue) Remove(ctx context.Context, id string) (string, error) {
	return s.repository.Remove(ctx, id)
}

func (s *product) InsertProduct(ctx context.Context, name, details, description, typ string, tags []string, price float64, discount float32) error {
	a := domain.NewProduct(name, details, description, typ, tags, price, discount)
	if err := s.repository.Insert(ctx, *a); err != nil {
		return err
	}

	return nil
}

func (s *product) UpdateProduct(ctx context.Context, id, name, details, description, typ string, tags []string, price float64, discount float32) error {
	a := domain.NewProduct(name, details, description, typ, tags, price, discount)
	if err := s.repository.Update(ctx, *a, id); err != nil {
		return err
	}

	return nil
}

func (s *product) FindOneProduct(ctx context.Context, id string) (*domain.Product, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *product) ListAllProducts(ctx context.Context, skip, take uint64) (*[]domain.Product, error) {
	res, err := s.repository.ListAllProducts(ctx, skip, take)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *product) QueryProduct(ctx context.Context, skip uint64, take uint64) ([]domain.Product, error) {
	return nil, nil
}

func (s *product) RemoveProduct(ctx context.Context, id string) (string, error) {
	return s.repository.Remove(ctx, id)
}
