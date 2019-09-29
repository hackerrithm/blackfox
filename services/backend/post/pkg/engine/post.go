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

	"github.com/hackerrithm/blackfox/services/backend/post/pkg/domain"
)

type (
	// Post ...
	Post interface {
		Insert(ctx context.Context, author, topic, category, contentText, contentPhoto string) error

		// Update is the update-a-post use-case
		Update(ctx context.Context, author, topic, category, contentText, contentPhoto string, id string) error

		// Query is the list-the-posts use-case
		Query(ctx context.Context, skip uint64, take uint64) ([]domain.Post, error)

		// FindOne ...
		FindOne(ctx context.Context, id string) (*domain.Post, error)

		// RemoveDelete ...
		Remove(ctx context.Context, id string) (string, error)

		// ListAllPosts ...
		ListAllPosts(ctx context.Context, skip, take uint64) (*[]domain.Post, error)
	}

	post struct {
		repository PostRepository
	}
)

var (
	postInstance Post
	postOnce     sync.Once
)

func (f *engineFactory) NewPost() Post {
	postOnce.Do(func() {
		postInstance = &post{
			repository: f.NewPostRepository(),
			// jwt:        f.JWTSignParser,
		}
	})
	return postInstance
}
