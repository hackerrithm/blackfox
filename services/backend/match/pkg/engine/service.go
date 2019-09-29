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

	"github.com/hackerrithm/blackfox/services/backend/match/pkg/domain"
)

func (s *match) Insert(ctx context.Context, person, details, description, typ string, similarities []string) error {
	a := domain.NewMatch(person, details, description, typ, similarities)
	if err := s.repository.Insert(ctx, *a); err != nil {
		return err
	}

	return nil
}

func (s *match) Update(ctx context.Context, id, person, details, description, typ string, similarities []string) error {
	a := domain.NewMatch(person, details, description, typ, similarities)
	if err := s.repository.Update(ctx, *a, id); err != nil {
		return err
	}

	return nil
}

func (s *match) FindOne(ctx context.Context, id string) (*domain.Match, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *match) ListAllMatches(ctx context.Context, skip, take uint64) (*[]domain.Match, error) {
	// search all users (limit) with rank > n
	// search all goals for n users
	// compare similarities (limit) to current user
	// where most frequent appearance return users "influencers"
	
	res, err := s.repository.ListAllMatches(ctx, skip, take)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *match) Query(ctx context.Context, skip uint64, take uint64) ([]domain.Match, error) {
	return nil, nil
}

func (s *match) Remove(ctx context.Context, id string) (string, error) {
	return s.repository.Remove(ctx, id)
}
