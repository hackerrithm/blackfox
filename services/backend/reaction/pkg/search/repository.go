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

package search

import (
	"context"

	"github.com/hackerrithm/blackfox/services/backend/reaction/pkg/domain"
)

// Repository ...
type Repository interface {
	Close()
	InsertReaction(ctx context.Context, reaction domain.Reaction) error
	SearchReactions(ctx context.Context, query string, skip uint64, take uint64) ([]domain.Reaction, error)
}

var impl Repository

// SetRepository ...
func SetRepository(repository Repository) {
	impl = repository
}

// Close ...
func Close() {
	impl.Close()
}

// InsertReaction ...
func InsertReaction(ctx context.Context, reaction domain.Reaction) error {
	return impl.InsertReaction(ctx, reaction)
}

// PerformSearchReactions ...
func PerformSearchReactions(ctx context.Context, query string, skip uint64, take uint64) ([]domain.Reaction, error) {
	return impl.SearchReactions(ctx, query, skip, take)
}
