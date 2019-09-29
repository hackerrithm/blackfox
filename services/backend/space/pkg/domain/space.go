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

var spaceContextKey contextKey = "space"

type contextKey string

// Space ...
type Space struct {
	ID          bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Creator     bson.ObjectId   `json:"author" bson:"author,omitempty"`
	Managers    []bson.ObjectId `json:"managers" bson:"managers,omitempty"`
	Topic       string          `json:"topic" bson:"topic,omitempty"`
	Details     string          `json:"details" bson:"details,omitempty"`
	Description string          `json:"description" bson:"description,omitempty"`
	Type        string          `json:"type" bson:"type,omitempty"`
	Tags        []string        `json:"tags" bson:"tags,omitempty"`
	Date        time.Time       `json:"timestamp" bson:"timestamp,omitempty"`
	Followers   []bson.ObjectId `json:"followers" bson:"followers,omitempty"`
}

// NewSpace creates a new Space!
func NewSpace(creator, topic, details, description, typ string, managers, followers, tags []string) *Space {
	var mngs, fllwrs []bson.ObjectId
	for _, m := range managers {
		mngs = append(mngs, bson.ObjectIdHex(m))
	}
	for _, f := range followers {
		fllwrs = append(fllwrs, bson.ObjectIdHex(f))
	}
	return &Space{
		Creator:     bson.ObjectIdHex(creator),
		Topic:       topic,
		Details:     details,
		Description: description,
		Type:        typ,
		Followers:   fllwrs,
		Managers:    mngs,
		Tags:        tags,
		Date:        now(),
	}
}
