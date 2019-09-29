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

var messageContextKey contextKey = "message"

type (
	// Message is a message object
	// that has datatypes for
	// the message domain (Message)
	Message struct {
		ID         bson.ObjectId `json:"_id" bson:"_id"`
		Sender     bson.ObjectId `json:"sender" bson:"sender"`
		Receiver   bson.ObjectId `json:"receiver" bson:"receiver"`
		Type       string        `json:"type" bson:"type"`
		Text       string        `json:"text" bson:"text"`
		IsSeen     bool          `json:"is_seen" bson:"is_seen"`
		IsSent     bool          `json:"is_sent" bson:"is_sent"`
		IsReceived bool          `json:"is_received" bson:"is_received"`
		Timestamp  time.Time     `json:"timestamp" bson:"timestamp"`
	}
)

// NewMessage creates a new Message!
func NewMessage(sender, receiver, msgType, text string, isSeen, isSent, isReceived bool) *Message {
	return &Message{
		Sender:     bson.ObjectIdHex(sender),
		Receiver:   bson.ObjectIdHex(receiver),
		Type:       msgType,
		Text:       text,
		IsSeen:     isSeen,
		IsSent:     isSent,
		IsReceived: isReceived,
		Timestamp:  now(),
	}
}
