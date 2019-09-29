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

	"github.com/hackerrithm/blackfox/services/backend/group/pkg/domain"
)

type (
	// Group ...
	Group interface {
		Insert(ctx context.Context, title, details, description, typ string, people []string) error

		// Update is the update-a-group use-case
		Update(ctx context.Context, id, title, details, description, typ string, people []string) error

		// Query is the list-the-groups use-case
		Query(ctx context.Context, skip uint64, take uint64) ([]domain.Group, error)

		// FindOne ...
		FindOne(ctx context.Context, id string) (*domain.Group, error)

		// RemoveDelete ...
		Remove(ctx context.Context, id string) (string, error)

		// ListAllGroups ...
		ListAllGroups(ctx context.Context, skip, take uint64) (*[]domain.Group, error)
	}

	group struct {
		repository GroupRepository
	}
)

var (
	groupInstance Group
	groupOnce     sync.Once
)

func (f *engineFactory) NewGroup() Group {
	groupOnce.Do(func() {
		groupInstance = &group{
			repository: f.NewGroupRepository(),
			// jwt:        f.JWTSignParser,
		}
	})
	return groupInstance
}
