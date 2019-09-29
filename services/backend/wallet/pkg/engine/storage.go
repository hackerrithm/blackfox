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

	"github.com/hackerrithm/blackfox/services/backend/wallet/pkg/domain"
)

type (
	// WalletRepository defines the methods that any
	// data storage provider needs to implement to get
	// and store wallets
	WalletRepository interface {

		// Wallet adds a new Wallet to the datastore
		Insert(c context.Context, wallet domain.Wallet) error

		// Put adds a new Wallet to the datastore
		Update(c context.Context, wallet domain.Wallet, id string) error

		// Query returns existing wallets matching the
		// query provided
		Query(c context.Context, query *Query) []*domain.Wallet

		// FindOne returns ...
		FindOne(c context.Context, id string) (*domain.Wallet, error)

		// Remove ...
		Remove(c context.Context, id string) (string, error)

		// ListAllWallets ...
		ListAllWallets(ctx context.Context, skip uint64, take uint64) ([]domain.Wallet, error)
	}

	// StorageFactory is the interface that a storage
	// provider needs to implement so that the engine can
	// request repository instances as it needs them
	StorageFactory interface {
		// NewWalletRepository returns a storage specific
		// WalletRepository implementation
		NewWalletRepository() WalletRepository

		Close()
	}
)
