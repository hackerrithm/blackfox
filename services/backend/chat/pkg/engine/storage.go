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

	"github.com/hackerrithm/blackfox/services/backend/chat/pkg/domain"
)

type (
	// ChatRepository defines the methods that any
	// data storage provider needs to implement to get
	// and store chats
	ChatRepository interface {

		// Chat adds a new Chat to the datastore
		Insert(c context.Context, chat domain.Chat) error

		// Put adds a new Chat to the datastore
		Update(c context.Context, chat domain.Chat, id string) error

		// Query returns existing chats matching the
		// query provided
		Query(c context.Context, query *Query) []*domain.Chat

		// FindOne returns ...
		FindOne(c context.Context, id string) (*domain.Chat, error)

		// Remove ...
		Remove(c context.Context, id string) (string, error)

		// ListAllChats ...
		ListAllChats(ctx context.Context, skip uint64, take uint64) ([]domain.Chat, error)
	}

	// StorageFactory is the interface that a storage
	// provider needs to implement so that the engine can
	// request repository instances as it needs them
	StorageFactory interface {
		// NewChatRepository returns a storage specific
		// ChatRepository implementation
		NewChatRepository() ChatRepository

		Close()
	}
)
