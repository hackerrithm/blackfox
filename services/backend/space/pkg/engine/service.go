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
	"log"

	"github.com/hackerrithm/blackfox/services/backend/space/pkg/domain"
)

func (s *space) Insert(ctx context.Context, creator, topic, details, description, typ string, managers, followers, tags []string) error {
	log.Println("got in service")
	a := domain.NewSpace(creator, topic, details, description, typ, managers, followers, tags)
	if err := s.repository.Insert(ctx, *a); err != nil {
		return err
	}

	return nil
}

func (s *space) Update(ctx context.Context, id uint64, creator, topic, details, description, typ string, managers, followers, tags []string) error {
	a := domain.NewSpace(creator, topic, details, description, typ, managers, followers, tags)
	if err := s.repository.Update(ctx, *a, id); err != nil {
		return err
	}

	return nil
}

func (s *space) FindOne(ctx context.Context, id uint64) (*domain.Space, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *space) ListAllSpaces(ctx context.Context, skip, take uint64) (*[]domain.Space, error) {
	res, err := s.repository.ListAllSpaces(ctx, skip, take)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *space) Query(ctx context.Context, skip uint64, take uint64) ([]domain.Space, error) {
	return nil, nil
}

func (s *space) Remove(ctx context.Context, id uint64) (string, error) {
	return s.repository.Remove(ctx, id)
}
