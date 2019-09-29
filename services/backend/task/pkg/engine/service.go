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

	"github.com/hackerrithm/blackfox/services/backend/task/pkg/domain"
)

func (s *task) Insert(ctx context.Context, text string) error {
	a := domain.NewTask(text)
	if err := s.repository.Insert(ctx, *a); err != nil {
		return err
	}

	return nil
}

func (s *task) Update(ctx context.Context, id, text string) error {
	a := domain.NewTask(text)
	if err := s.repository.Update(ctx, *a, id); err != nil {
		return err
	}

	return nil
}

func (s *task) FindOne(ctx context.Context, id string) (*domain.Task, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *task) ListAllTasks(ctx context.Context, skip, take uint64) (*[]domain.Task, error) {
	res, err := s.repository.ListAllTasks(ctx, skip, take)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *task) Query(ctx context.Context, skip uint64, take uint64) ([]domain.Task, error) {
	return nil, nil
}

func (s *task) Remove(ctx context.Context, id string) (string, error) {
	return s.repository.Remove(ctx, id)
}
