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

	"github.com/hackerrithm/blackfox/services/backend/reaction/pkg/domain"
)

type (
	// ReactionRepository defines the methods that any
	// data storage provider needs to implement to get
	// and store reactions
	ReactionRepository interface {

		// Reaction adds a new Reaction to the datastore
		Insert(c context.Context, reaction domain.Reaction) error

		// Put adds a new Reaction to the datastore
		Update(c context.Context, reaction domain.Reaction, id string) error

		// Query returns existing reactions reactioning the
		// query provided
		Query(c context.Context, query *Query) []*domain.Reaction

		// FindOne returns ...
		FindOne(c context.Context, id string) (*domain.Reaction, error)

		// Remove ...
		Remove(c context.Context, id string) (string, error)

		// ListAllReactiones ...
		ListAllReactiones(ctx context.Context, skip uint64, take uint64) ([]domain.Reaction, error)
	}

	// StorageFactory is the interface that a storage
	// provider needs to implement so that the engine can
	// request repository instances as it needs them
	StorageFactory interface {
		// NewReactionRepository returns a storage specific
		// ReactionRepository implementation
		NewReactionRepository() ReactionRepository

		Close()
	}
)
