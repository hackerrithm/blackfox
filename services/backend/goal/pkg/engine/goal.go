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

	"github.com/hackerrithm/blackfox/services/backend/goal/pkg/domain"
)

type (
	// Goal ...
	Goal interface {
		Insert(ctx context.Context, creator, aim, reason, details, typ string, tags []string) error

		// Update is the update-a-goal use-case
		Update(ctx context.Context, id, creator, aim, reason, details, typ string, tags []string) error

		// Query is the list-the-goals use-case
		Query(ctx context.Context, skip uint64, take uint64) ([]domain.Goal, error)

		// FindOne ...
		FindOne(ctx context.Context, id string) (*domain.Goal, error)

		// RemoveDelete ...
		Remove(ctx context.Context, id string) (string, error)

		// ListAllGoals ...
		ListAllGoals(ctx context.Context, skip, take uint64) (*[]domain.Goal, error)
	}

	goal struct {
		repository GoalRepository
	}
)

var (
	goalInstance Goal
	goalOnce     sync.Once
)

func (f *engineFactory) NewGoal() Goal {
	goalOnce.Do(func() {
		goalInstance = &goal{
			repository: f.NewGoalRepository(),
			// jwt:        f.JWTSignParser,
		}
	})
	return goalInstance
}
