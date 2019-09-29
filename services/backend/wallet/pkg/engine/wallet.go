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

	"github.com/hackerrithm/blackfox/services/backend/wallet/pkg/domain"
)

type (
	// Wallet ...
	Wallet interface {
		Insert(ctx context.Context, userID, details, description, typ, currencyAbbr, currencyName, currencyType, tokenType string,
			balanceAmount float64, tokenVal float32, tokenAmt int64) error

		// Update is the update-a-wallet use-case
		Update(ctx context.Context, id, userID, details, description, typ, currencyAbbr, currencyName, currencyType, tokenType string,
			balanceAmount float64, tokenVal float32, tokenAmt int64) error

		// Query is the list-the-wallets use-case
		Query(ctx context.Context, skip uint64, take uint64) ([]domain.Wallet, error)

		// FindOne ...
		FindOne(ctx context.Context, id string) (*domain.Wallet, error)

		// RemoveDelete ...
		Remove(ctx context.Context, id string) (string, error)

		// ListAllWallets ...
		ListAllWallets(ctx context.Context, skip, take uint64) (*[]domain.Wallet, error)
	}

	wallet struct {
		repository WalletRepository
	}
)

var (
	walletInstance Wallet
	walletOnce     sync.Once
)

func (f *engineFactory) NewWallet() Wallet {
	walletOnce.Do(func() {
		walletInstance = &wallet{
			repository: f.NewWalletRepository(),
			// jwt:        f.JWTSignParser,
		}
	})
	return walletInstance
}
