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

var groupContextKey contextKey = "group"

type contextKey string

// Group ...
type Group struct {
	ID          bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Title       string          `json:"title" bson:"title,omitempty"`
	People      []bson.ObjectId `json:"people" bson:"people,omitempty"`
	Details     string          `json:"details" bson:"details,omitempty"`
	Description string          `json:"description" bson:"description,omitempty"`
	Type        string          `json:"type" bson:"type,omitempty"`
	Date        time.Time       `json:"timestamp" bson:"timestamp,omitempty"`
}

// NewGroup creates a new Group!
func NewGroup(title, details, description, typ string, people []string) *Group {
	var ppl []bson.ObjectId

	for _, p := range people {
		ppl = append(ppl, bson.ObjectIdHex(p))
	}
	return &Group{
		Title:       title,
		People:      ppl,
		Details:     details,
		Description: description,
		Type:        typ,
		Date:        now(),
	}
}
