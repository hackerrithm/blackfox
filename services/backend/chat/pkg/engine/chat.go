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

	"github.com/hackerrithm/blackfox/services/backend/chat/pkg/domain"
)

type (
	// Chat ...
	Chat interface {
		Insert(ctx context.Context, creator, topic, details, description, typ string, managers, followers, tags []string) error

		// Update is the update-a-chat use-case
		Update(ctx context.Context, id, creator, topic, details, description, typ string, managers, followers, tags []string) error

		// Query is the list-the-chats use-case
		Query(ctx context.Context, skip uint64, take uint64) ([]domain.Chat, error)

		// FindOne ...
		FindOne(ctx context.Context, id string) (*domain.Chat, error)

		// RemoveDelete ...
		Remove(ctx context.Context, id string) (string, error)

		// ListAllChats ...
		ListAllChats(ctx context.Context, skip, take uint64) (*[]domain.Chat, error)
	}

	chat struct {
		repository ChatRepository
	}
)

var (
	chatInstance Chat
	chatOnce     sync.Once
)

func (f *engineFactory) NewChat() Chat {
	chatOnce.Do(func() {
		chatInstance = &chat{
			repository: f.NewChatRepository(),
			// jwt:        f.JWTSignParser,
		}
	})
	return chatInstance
}
