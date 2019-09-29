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

	"github.com/hackerrithm/blackfox/services/backend/group/pkg/domain"
)

func (s *group) Insert(ctx context.Context, title, details, description, typ string, people []string) error {
	a := domain.NewGroup(title, details, description, typ, people)
	if err := s.repository.Insert(ctx, *a); err != nil {
		return err
	}

	return nil
}

func (s *group) Update(ctx context.Context, id, title, details, description, typ string, people []string) error {
	a := domain.NewGroup(title, details, description, typ, people)
	if err := s.repository.Update(ctx, *a, id); err != nil {
		return err
	}

	return nil
}

func (s *group) FindOne(ctx context.Context, id string) (*domain.Group, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *group) ListAllGroups(ctx context.Context, skip, take uint64) (*[]domain.Group, error) {
	res, err := s.repository.ListAllGroups(ctx, skip, take)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *group) Query(ctx context.Context, skip uint64, take uint64) ([]domain.Group, error) {
	return nil, nil
}

func (s *group) Remove(ctx context.Context, id string) (string, error) {
	return s.repository.Remove(ctx, id)
}
