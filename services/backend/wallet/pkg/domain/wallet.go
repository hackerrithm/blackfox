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

package domain

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

var walletContextKey contextKey = "wallet"

type contextKey string

// Wallet ...
type Wallet struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserID       bson.ObjectId `json:"user_id" bson:"user_id,omitempty"`
	Details      string        `json:"details" bson:"details,omitempty"`
	Description  string        `json:"description" bson:"description,omitempty"`
	Type         string        `json:"type" bson:"type,omitempty"`
	Tokens       []Token       `json:"tokens" bson:"tokens,omitempty"`
	Balance      Money         `json:"balance" bson:"balance,omitempty"`
	CreatedOn    time.Time     `json:"created_on" bson:"createdOn,omitempty"`
	ExpiresOn    time.Time     `json:"expires_on" bson:"expiresOn,omitempty"`
	LastAccessed time.Time     `json:"last_accessed" bson:"lastAccessed,omitempty"`
}

// Currency ...
type Currency struct {
	Abbreviation string `json:"abbreviation" bson:"abbreviation,omitempty"`
	Name         string `json:"name" bson:"name,omitempty"`
	Type         string `json:"type" bson:"type,omitempty"`
}

// Token ...
type Token struct {
	Type   string  `json:"type" bson:"type,omitempty"`
	Amount int64   `json:"amount" bson:"amount,omitempty"`
	Value  float32 `json:"value" bson:"value,omitempty"`
}

// Money ...
type Money struct {
	Currency Currency `json:"currency" bson:"currency,omitempty"`
	Amount   float64  `json:"amount" bson:"amount,omitempty"`
}

// NewWallet creates a new Wallet!
func NewWallet(userID, details, description, typ, currencyAbbr, currencyName, currencyType, tokenType string,
	balanceAmount float64, tokenVal float32, tokenAmt int64) *Wallet {
	return &Wallet{
		UserID:      bson.ObjectIdHex(userID),
		Details:     details,
		Description: description,
		Type:        typ,
		Balance: Money{
			Currency: Currency{
				Abbreviation: currencyAbbr,
				Name:         currencyName,
				Type:         currencyType,
			},
			Amount: balanceAmount,
		},
		Tokens:       make([]Token, 0),
		CreatedOn:    now(),
		ExpiresOn:    now(),
		LastAccessed: now(),
	}
}
