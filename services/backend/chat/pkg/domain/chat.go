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

var chatContextKey contextKey = "chat"

type (
	contextKey string

	// Chat is a conversation object
	// that has datatypes for
	// the chat domain (Chat)
	Chat struct {
		ID        bson.ObjectId            `json:"_id" bson:"_id"`
		Messages  []Message                `json:"messages" bson:"messages"`
		Observers map[string]chan *Message `json:"observers" bson:"observers"`
		StartDate time.Time                `json:"startdate" bson:"startdate"`
		EndDate   time.Time                `json:"enddate" bson:"enddate"`
	}
)

// NewChat creates a new Chat!
func NewChat(messages []Message) *Chat {
	return &Chat{
		Messages:  messages,
		StartDate: now(),
	}
}
