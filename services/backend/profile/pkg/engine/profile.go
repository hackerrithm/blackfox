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

	"github.com/hackerrithm/blackfox/services/backend/profile/pkg/domain"
)

type (
	// Profile ...
	Profile interface {
		Insert(ctx context.Context, username, level, about string, followers, following []string, rings int32) error

		// Update is the update-a-profile use-case
		Update(ctx context.Context, id, username, level, about string, followers, following []string, rings int32) error

		// Query is the list-the-profiles use-case
		Query(ctx context.Context, skip uint64, take uint64) ([]domain.Profile, error)

		// FindOne ...
		FindOne(ctx context.Context, id string) (*domain.Profile, error)

		// RemoveDelete ...
		Remove(ctx context.Context, id string) (string, error)

		// ListAllProfiles ...
		ListAllProfiles(ctx context.Context, skip, take uint64) (*[]domain.Profile, error)
	}

	profile struct {
		repository ProfileRepository
	}
)

var (
	profileInstance Profile
	profileOnce     sync.Once
)

func (f *engineFactory) NewProfile() Profile {
	profileOnce.Do(func() {
		profileInstance = &profile{
			repository: f.NewProfileRepository(),
			// jwt:        f.JWTSignParser,
		}
	})
	return profileInstance
}
