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

	"github.com/hackerrithm/blackfox/services/backend/reaction/pkg/domain"
)

type (
	// Reaction ...
	Reaction interface {
		Insert(ctx context.Context, id, person, details, description, typ string) error

		// Update is the update-a-reaction use-case
		Update(ctx context.Context, id, person, details, description, typ string) error

		// Query is the list-the-reactions use-case
		Query(ctx context.Context, skip uint64, take uint64) ([]domain.Reaction, error)

		// FindOne ...
		FindOne(ctx context.Context, id string) (*domain.Reaction, error)

		// RemoveDelete ...
		Remove(ctx context.Context, id string) (string, error)

		// ListAllReactiones ...
		ListAllReactiones(ctx context.Context, skip, take uint64) (*[]domain.Reaction, error)
	}

	reaction struct {
		repository ReactionRepository
	}
)

var (
	reactionInstance Reaction
	reactionOnce     sync.Once
)

func (f *engineFactory) NewReaction() Reaction {
	reactionOnce.Do(func() {
		reactionInstance = &reaction{
			repository: f.NewReactionRepository(),
			// jwt:        f.JWTSignParser,
		}
	})
	return reactionInstance
}
