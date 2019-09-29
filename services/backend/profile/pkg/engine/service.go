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

	"github.com/hackerrithm/blackfox/services/backend/profile/pkg/domain"
)

func (s *profile) Insert(ctx context.Context, username, level, about string, followers, following []string, rings int32) error {
	a := domain.NewProfile(username, level, about, followers, following, rings)
	if err := s.repository.Insert(ctx, *a); err != nil {
		return err
	}

	return nil
}

func (s *profile) Update(ctx context.Context, id, username, level, about string, followers, following []string, rings int32) error {
	a := domain.NewProfile(username, level, about, followers, following, rings)
	if err := s.repository.Update(ctx, *a, id); err != nil {
		return err
	}

	return nil
}

func (s *profile) FindOne(ctx context.Context, id string) (*domain.Profile, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *profile) ListAllProfiles(ctx context.Context, skip, take uint64) (*[]domain.Profile, error) {
	res, err := s.repository.ListAllProfiles(ctx, skip, take)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *profile) Query(ctx context.Context, skip uint64, take uint64) ([]domain.Profile, error) {
	return nil, nil
}

func (s *profile) Remove(ctx context.Context, id string) (string, error) {
	return s.repository.Remove(ctx, id)
}
