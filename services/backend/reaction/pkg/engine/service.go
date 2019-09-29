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

	"github.com/hackerrithm/blackfox/services/backend/reaction/pkg/domain"
)

func (s *reaction) Insert(ctx context.Context, person, details, description, typ string) error {
	a := domain.NewReaction(person, details, description, typ)
	if err := s.repository.Insert(ctx, *a); err != nil {
		return err
	}

	return nil
}

func (s *reaction) Update(ctx context.Context, id, person, details, description, typ string) error {
	a := domain.NewReaction(person, details, description, typ)
	if err := s.repository.Update(ctx, *a, id); err != nil {
		return err
	}

	return nil
}

func (s *reaction) FindOne(ctx context.Context, id string) (*domain.Reaction, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *reaction) ListAllReactiones(ctx context.Context, skip, take uint64) (*[]domain.Reaction, error) {
	res, err := s.repository.ListAllReactiones(ctx, skip, take)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *reaction) Query(ctx context.Context, skip uint64, take uint64) ([]domain.Reaction, error) {
	return nil, nil
}

func (s *reaction) Remove(ctx context.Context, id string) (string, error) {
	return s.repository.Remove(ctx, id)
}
