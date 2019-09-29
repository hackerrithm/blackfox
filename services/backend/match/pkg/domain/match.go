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

var matchContextKey contextKey = "match"

type contextKey string

// Match ...
type Match struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	PersonID     bson.ObjectId `json:"person_id" bson:"person_id,omitempty"`
	Details      string        `json:"details" bson:"details,omitempty"`
	Description  string        `json:"description" bson:"description,omitempty"`
	Type         string        `json:"type" bson:"type,omitempty"`
	Similarities []string      `json:"similarities" bson:"similarities,omitempty"`
	Date         time.Time     `json:"timestamp" bson:"timestamp,omitempty"`
}

// MatchedUser ...
type MatchedUser struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name         string        `json:"name" bson:"name,omitempty"`
	Username     string        `json:"username" bson:"username,omitempty"`
	Firstname    string        `json:"firstname" bson:"firstname,omitempty"`
	Lastname     string        `json:"lastname" bson:"lastname,omitempty"`
	Status       string        `json:"status" bson:"status,omitempty"`
	Type         string        `json:"account_type" bson:"account_type,omitempty"`
	EmailAddress string        `json:"emailaddress" bson:"email_address,omitempty"`
	Gender       string        `json:"gender" bson:"gender,omitempty"`
}

// NewMatch creates a new Match!
func NewMatch(person, details, description, typ string, similarities []string) *Match {
	return &Match{
		PersonID:     bson.ObjectId(person),
		Details:      details,
		Description:  description,
		Type:         typ,
		Similarities: similarities,
		Date:         now(),
	}
}
