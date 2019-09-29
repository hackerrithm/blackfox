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

var orderContextKey contextKey = "order"

type contextKey string

type (
	// Order ...
	Order struct {
		ID          bson.ObjectId    `json:"id" bson:"_id,omitempty"`
		UserID      bson.ObjectId    `json:"userID" bson:"user_id,omitempty"`
		Details     string           `json:"details" bson:"details,omitempty"`
		Description string           `json:"description" bson:"description,omitempty"`
		Type        string           `json:"type" bson:"type,omitempty"`
		TotalPrice  float64          `json:"totalPrice" bson:"total_price,omitempty"`
		Products    []OrderedProduct `json:"products" bson:"products,omitempty"`
		CreatedOn   time.Time        `json:"createdOn" bson:"created_on,omitempty"`
	}

	// OrderedProduct ...
	OrderedProduct struct {
		ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name        string        `json:"name" bson:"name,omitempty"`
		Description string        `json:"description" bson:"description,omitempty"`
		Price       float64       `json:"price" bson:"price,omitempty"`
		Quantity    uint32        `json:"quantity" bson:"quantity,omitempty"`
	}
)

// NewOrder creates a new Order!
func NewOrder(userID, details, description, typ string, totalPrice float64, products []OrderedProduct) *Order {
	return &Order{
		UserID:      bson.ObjectIdHex(userID),
		Details:     details,
		Description: description,
		Type:        typ,
		TotalPrice:  totalPrice,
		Products:    products,
		CreatedOn:   now(),
	}
}
