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
	"context"

	"github.com/hackerrithm/blackfox/services/backend/wallet/pkg/domain"
)

func (s *wallet) Insert(ctx context.Context, userID, details, description, typ, currencyAbbr, currencyName, currencyType, tokenType string,
	balanceAmount float64, tokenVal float32, tokenAmt int64) error {
	a := domain.NewWallet(userID, details, description, typ, currencyAbbr, currencyName, currencyType, tokenType,
		balanceAmount, tokenVal, tokenAmt)
	if err := s.repository.Insert(ctx, *a); err != nil {
		return err
	}

	return nil
}

func (s *wallet) Update(ctx context.Context, id, userID, details, description, typ, currencyAbbr, currencyName, currencyType, tokenType string,
	balanceAmount float64, tokenVal float32, tokenAmt int64) error {
	a := domain.NewWallet(userID, details, description, typ, currencyAbbr, currencyName, currencyType, tokenType,
		balanceAmount, tokenVal, tokenAmt)
	if err := s.repository.Update(ctx, *a, id); err != nil {
		return err
	}

	return nil
}

func (s *wallet) FindOne(ctx context.Context, id string) (*domain.Wallet, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *wallet) ListAllWallets(ctx context.Context, skip, take uint64) (*[]domain.Wallet, error) {
	res, err := s.repository.ListAllWallets(ctx, skip, take)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *wallet) Query(ctx context.Context, skip uint64, take uint64) ([]domain.Wallet, error) {
	return nil, nil
}

func (s *wallet) Remove(ctx context.Context, id string) (string, error) {
	return s.repository.Remove(ctx, id)
}
