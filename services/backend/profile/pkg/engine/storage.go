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
	"golang.org/x/net/context"

	"github.com/hackerrithm/blackfox/services/backend/profile/pkg/domain"
)

type (
	// ProfileRepository defines the methods that any
	// data storage provider needs to implement to get
	// and store profiles
	ProfileRepository interface {

		// Profile adds a new Profile to the datastore
		Insert(c context.Context, profile domain.Profile) error

		// Put adds a new Profile to the datastore
		Update(c context.Context, profile domain.Profile, id string) error

		// Query returns existing profiles matching the
		// query provided
		Query(c context.Context, query *Query) []*domain.Profile

		// FindOne returns ...
		FindOne(c context.Context, id string) (*domain.Profile, error)

		// Remove ...
		Remove(c context.Context, id string) (string, error)

		// ListAllProfiles ...
		ListAllProfiles(ctx context.Context, skip uint64, take uint64) ([]domain.Profile, error)
	}

	// StorageFactory is the interface that a storage
	// provider needs to implement so that the engine can
	// request repository instances as it needs them
	StorageFactory interface {
		// NewProfileRepository returns a storage specific
		// ProfileRepository implementation
		NewProfileRepository() ProfileRepository

		Close()
	}
)
