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

	"github.com/hackerrithm/blackfox/services/backend/space/pkg/domain"
)

type (
	// Space ...
	Space interface {
		Insert(ctx context.Context, creator, topic, details, description, typ string, managers, followers, tags []string) error

		// Update is the update-a-space use-case
		Update(ctx context.Context, id, creator, topic, details, description, typ string, managers, followers, tags []string) error

		// Query is the list-the-spaces use-case
		Query(ctx context.Context, skip uint64, take uint64) ([]domain.Space, error)

		// FindOne ...
		FindOne(ctx context.Context, id string) (*domain.Space, error)

		// RemoveDelete ...
		Remove(ctx context.Context, id string) (string, error)

		// ListAllSpaces ...
		ListAllSpaces(ctx context.Context, skip, take uint64) (*[]domain.Space, error)
	}

	space struct {
		repository SpaceRepository
	}
)

var (
	spaceInstance Space
	spaceOnce     sync.Once
)

func (f *engineFactory) NewSpace() Space {
	spaceOnce.Do(func() {
		spaceInstance = &space{
			repository: f.NewSpaceRepository(),
			// jwt:        f.JWTSignParser,
		}
	})
	return spaceInstance
}
