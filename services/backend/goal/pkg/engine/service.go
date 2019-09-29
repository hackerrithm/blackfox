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

	"github.com/hackerrithm/blackfox/services/backend/goal/pkg/domain"
)

func (s *goal) Insert(ctx context.Context, creator, aim, reason, details, typ string, tags []string) error {
	log.Println("arrived at service")
	a := domain.NewGoal(creator, aim, reason, details, typ, tags)
	if err := s.repository.Insert(ctx, *a, domain.Journey{}); err != nil {
		return err
	}

	return nil
}

func (s *goal) Update(ctx context.Context, id, creator, aim, reason, details, typ string, tags []string) error {
	a := domain.NewGoal(creator, aim, reason, details, typ, tags)
	if err := s.repository.Update(ctx, *a, domain.Journey{}, id); err != nil {
		return err
	}

	return nil
}

func (s *goal) FindOne(ctx context.Context, id string) (*domain.Goal, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *goal) ListAllGoals(ctx context.Context, skip, take uint64) (*[]domain.Goal, error) {
	res, err := s.repository.ListAllGoals(ctx, skip, take)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *goal) Query(ctx context.Context, skip uint64, take uint64) ([]domain.Goal, error) {
	return nil, nil
}

func (s *goal) Remove(ctx context.Context, id string) (string, error) {
	return s.repository.Remove(ctx, id)
}
