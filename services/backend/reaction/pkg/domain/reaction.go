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

var reactionContextKey contextKey = "reaction"

type contextKey string

// Reaction ...
type Reaction struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	PersonID    bson.ObjectId `json:"person_id" bson:"person_id,omitempty"`
	Details     string        `json:"details" bson:"details,omitempty"`
	Description string        `json:"description" bson:"description,omitempty"`
	Type        string        `json:"type" bson:"type,omitempty"`
	Date        time.Time     `json:"timestamp" bson:"timestamp,omitempty"`
}

// NewReaction creates a new Reaction!
func NewReaction(person, details, description, typ string) *Reaction {
	return &Reaction{
		PersonID:    bson.ObjectId(person),
		Details:     details,
		Description: description,
		Type:        typ,
		Date:        now(),
	}
}
