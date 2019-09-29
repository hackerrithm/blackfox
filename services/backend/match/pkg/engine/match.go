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

	"github.com/hackerrithm/blackfox/services/backend/match/pkg/domain"
)

type (
	// Match ...
	Match interface {
		Insert(ctx context.Context, person, details, description, typ string, similarities []string) error

		// Update is the update-a-match use-case
		Update(ctx context.Context, id, person, details, description, typ string, similarities []string) error

		// Query is the list-the-matchs use-case
		Query(ctx context.Context, skip uint64, take uint64) ([]domain.Match, error)

		// FindOne ...
		FindOne(ctx context.Context, id string) (*domain.Match, error)

		// RemoveDelete ...
		Remove(ctx context.Context, id string) (string, error)

		// ListAllMatches ...
		ListAllMatches(ctx context.Context, skip, take uint64) (*[]domain.Match, error)
	}

	match struct {
		repository MatchRepository
	}
)

var (
	matchInstance Match
	matchOnce     sync.Once
)

func (f *engineFactory) NewMatch() Match {
	matchOnce.Do(func() {
		matchInstance = &match{
			repository: f.NewMatchRepository(),
			// jwt:        f.JWTSignParser,
		}
	})
	return matchInstance
}
