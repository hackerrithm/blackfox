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

package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"

	"github.com/hackerrithm/blackfox/services/backend/wallet/pkg/domain"
	pb "github.com/hackerrithm/blackfox/services/backend/wallet/pkg/model"
)

// Client ...
type Client struct {
	conn    *grpc.ClientConn
	service pb.WalletServiceClient
}

// NewClient ...
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second*100))
	if err != nil {
		return nil, err
	}
	c := pb.NewWalletServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close ...
func (c *Client) Close() {
	c.conn.Close()
}

// Post ...
func (c *Client) Post(ctx context.Context, userID, details, description, typ, currencyAbbr, currencyName, currencyType, tokenType string,
	balanceAmount float64, tokenVal float32, tokenAmt int64) (string, error) {
	r, err := c.service.PostWallet(
		ctx,
		&pb.PostWalletRequest{
			UserID: userID,
			Balance: &pb.Money{
				Amount: balanceAmount,
				Currency: &pb.Currency{
					Abbreviation: currencyAbbr,
					Name:         currencyName,
					Type:         currencyType,
				},
			},
			Details:     details,
			Description: description,
			Type:        typ,
			Tokens: &pb.Token{
				Amount: tokenAmt,
				Type:   tokenType,
				Value:  tokenVal,
			},
		},
	)
	if err != nil {
		return "", err
	}

	log.Println(r.Wallet)
	return r.Wallet, nil
}

// Put ...
func (c *Client) Put(ctx context.Context, id, userID, details, description, typ, currencyAbbr, currencyName, currencyType, tokenType string,
	balanceAmount float64, tokenVal float32, tokenAmt int64) (string, error) {
	r, err := c.service.PutWallet(
		ctx,
		&pb.PutWalletRequest{
			Id: id,
			Balance: &pb.Money{
				Amount: balanceAmount,
				Currency: &pb.Currency{
					Abbreviation: currencyAbbr,
					Name:         currencyName,
					Type:         currencyType,
				},
			},
			Details:     details,
			Description: description,
			Type:        typ,
			Tokens: &pb.Token{
				Amount: tokenAmt,
				Type:   tokenType,
				Value:  tokenVal,
			},
		},
	)
	if err != nil {
		return "", err
	}
	return r.Wallet, nil
}

// Get ...
func (c *Client) Get(ctx context.Context, id string, userID uint64) (*domain.Wallet, error) {
	r, err := c.service.GetWallet(
		ctx,
		&pb.GetWalletRequest{Id: id, UserID: userID},
	)
	if err != nil {
		return nil, err
	}

	return &domain.Wallet{
		ID: bson.ObjectIdHex(r.Wallet.Id),
		Balance: domain.Money{
			Amount: 0,
			Currency: domain.Currency{
				Abbreviation: r.Wallet.Balance.Currency.Abbreviation,
				Name:         r.Wallet.Balance.Currency.Name,
				Type:         r.Wallet.Balance.Currency.Type,
			},
		},
		Details:     r.Wallet.Details,
		Description: r.Wallet.Description,
		Type:        r.Wallet.Type,
		Tokens: domain.Token{
			Amount: r.Wallet.Tokens.Amount,
			Type:   r.Wallet.Tokens.Type,
			Value:  r.Wallet.Tokens.Value,
		},
	}, nil
}

// GetMultiple is used to get the list of specified wallets
func (c *Client) GetMultiple(ctx context.Context, skip uint64, take uint64) ([]domain.Wallet, error) {
	r, err := c.service.GetMultipleWallets(
		ctx,
		&pb.GetMultipleWalletsRequest{
			Skip: skip,
			Take: take,
		},
	)
	if err != nil {
		return nil, err
	}

	wallets := []domain.Wallet{}
	for _, a := range r.Wallets {
		wallets = append(wallets, domain.Wallet{
			ID: bson.ObjectIdHex(a.Id),
			Balance: domain.Money{
				Amount: 0,
				Currency: domain.Currency{
					Abbreviation: a.Balance.Currency.Abbreviation,
					Name:         a.Balance.Currency.Name,
					Type:         a.Balance.Currency.Type,
				},
			},
			Details:     a.Details,
			Description: a.Description,
			Type:        a.Type,
			Tokens: domain.Token{
				Amount: a.Tokens.Amount,
				Type:   a.Tokens.Type,
				Value:  a.Tokens.Value,
			},
		})
	}
	return wallets, nil
}

// Delete removes a wallet with passed identifier
func (c *Client) Delete(ctx context.Context, id string) (string, error) {
	r, err := c.service.DeleteWallet(
		ctx,
		&pb.DeleteWalletRequest{Id: id},
	)
	if err != nil {
		return "", err
	}
	return r.Id, nil
}
