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

var catalogueContextKey contextKey = "catalogue"

type contextKey string

type (
	// Catalogue ...
	Catalogue struct {
		ID          bson.ObjectId   `json:"id" bson:"_id,omitempty"`
		Products    []bson.ObjectId `json:"author" bson:"author,omitempty"`
		Name        string          `json:"name" bson:"name,omitempty"`
		Details     string          `json:"details" bson:"details,omitempty"`
		Description string          `json:"description" bson:"description,omitempty"`
		Type        string          `json:"type" bson:"type,omitempty"`
		Tags        []string        `json:"tags" bson:"tags,omitempty"`
		Date        time.Time       `json:"timestamp" bson:"timestamp,omitempty"`
	}

	// Product ...
	Product struct {
		ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name        string        `json:"name" bson:"name,omitempty"`
		Details     string        `json:"details" bson:"details,omitempty"`
		Description string        `json:"description" bson:"description,omitempty"`
		Images      []Image       `json:"productImage" bson:"product_image,omitempty"`
		Type        string        `json:"type" bson:"type,omitempty"`
		Price       float64       `json:"price" bson:"price,omitempty"`
		Discount    float32       `json:"discount" bson:"discount,omitempty"`
		Tags        []string      `json:"tags" bson:"tags,omitempty"`
		Date        time.Time     `json:"timestamp" bson:"timestamp,omitempty"`
	}

	// Image ...
	Image struct {
		Name   string `json:"name" bson:"name,omitempty"`
		Type   string `json:"type" bson:"type,omitempty"`
		Size   int64  `json:"size" bson:"size,omitempty"`
		Width  int    `json:"width" bson:"width,omitempty"`
		Height int    `json:"height" bson:"height,omitempty"`
	}
)

// NewCatalogue creates a new Catalogue!
func NewCatalogue(name, details, description, typ string, productIDs, tags []string) *Catalogue {
	var products []bson.ObjectId
	for _, p := range productIDs {
		products = append(products, bson.ObjectIdHex(p))
	}
	return &Catalogue{
		Name:        name,
		Details:     details,
		Description: description,
		Type:        typ,
		Products:    products,
		Tags:        tags,
		Date:        now(),
	}
}

// NewProduct creates a new Product!
func NewProduct(name, details, description, typ string, tags []string, price float64, discount float32) *Product {
	return &Product{
		Name:        name,
		Details:     details,
		Description: description,
		Type:        typ,
		Price:       price,
		Discount:    discount,
		Tags:        tags,
		Date:        now(),
	}
}
