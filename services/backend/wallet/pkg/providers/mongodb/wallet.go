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

package mongodb

import (
	"log"

	"golang.org/x/net/context"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	cfg "github.com/hackerrithm/blackfox/services/backend/wallet/configs"
	"github.com/hackerrithm/blackfox/services/backend/wallet/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/wallet/pkg/engine"
)

type (
	walletRepository struct {
		session *mgo.Session
	}
)

const (
	walletCollection = "wallet"
)

var (
	config cfg.Config
)

func newWalletRepository(session *mgo.Session) engine.WalletRepository {
	return &walletRepository{session}
}

func (r walletRepository) Insert(c context.Context, p domain.Wallet) error {
	s := r.session.Clone()
	defer s.Close()

	var wallet domain.Wallet

	wallet.UserID = p.UserID
	wallet.Description = p.Description
	wallet.Details = p.Details
	wallet.Balance = p.Balance
	wallet.Tokens = p.Tokens
	wallet.CreatedOn = p.CreatedOn
	wallet.ExpiresOn = p.ExpiresOn
	wallet.Type = p.Type
	wallet.LastAccessed = p.LastAccessed

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Insert(&wallet)
	if err != nil {
		log.Println("insert error ", err)
	}

	return nil
}

func (r walletRepository) Update(c context.Context, p domain.Wallet, id string) error {
	s := r.session.Clone()
	defer s.Close()

	var wallet domain.Wallet

	wallet.UserID = p.UserID
	wallet.Description = p.Description
	wallet.Details = p.Details
	wallet.Balance = p.Balance
	wallet.Tokens = p.Tokens
	wallet.CreatedOn = p.CreatedOn
	wallet.ExpiresOn = p.ExpiresOn
	wallet.Type = p.Type
	wallet.LastAccessed = p.LastAccessed

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Update(bson.M{"_id": toBSON(id)}, p)
	if err != nil {
		return err
	}

	return nil
}

func (r walletRepository) Query(c context.Context, query *engine.Query) []*domain.Wallet {
	s := r.session.Clone()
	defer s.Close()

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	p := []*domain.Wallet{}
	q := translateQuery(col, query)
	q.All(&p)

	return nil
}

func (r walletRepository) FindOne(c context.Context, id string) (*domain.Wallet, error) {
	s := r.session.Clone()
	defer s.Close()

	var wallet *domain.Wallet
	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"_id": toBSON(id)}).One(&wallet)

	if err != nil {
		return nil, nil
	}

	return wallet, nil
}

// ListAllWallets used for finding all user wallets
// by the passed skip and take parameters
func (r walletRepository) ListAllWallets(ctx context.Context, skip uint64, take uint64) ([]domain.Wallet, error) {
	s := r.session.Clone()
	defer s.Close()
	log.Println("got in repo")

	var wallets []domain.Wallet

	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(nil).All(&wallets)
	if err != nil {
		return nil, err
	}

	return wallets, nil
}

func (r walletRepository) Remove(c context.Context, id string) (string, error) {
	s := r.session.Clone()
	defer s.Close()

	err := s.DB(config.MongoDB).C(config.MongoCollection).Remove(bson.M{"_id": toBSON(id)}) //.One(&wallet)

	if err != nil {
		return "", nil
	}

	return "", nil
}
