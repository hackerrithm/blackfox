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

var postContextKey contextKey = "post"

type contextKey string

// Post ...
type Post struct {
	ID           bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Author       bson.ObjectId   `json:"author" bson:"author,omitempty"`
	Anonymous    bool            `json:"anonymous" bson:"anonymous,omitempty"`
	Topic        string          `json:"topic" bson:"topic,omitempty"`
	Category     string          `json:"category" bson:"category,omitempty"`
	ContentText  string          `json:"contentText" bson:"content_text,omitempty"`
	Type         string          `json:"type" bson:"type,omitempty"`
	Latitude     float64         `json:"latitude" bson:"latitude,omitempty"`
	Longitude    float64         `json:"longitude" bson:"longitude,omitempty"`
	Date         time.Time       `json:"timestamp" bson:"timestamp,omitempty"`
	ContentPhoto Image           `json:"contentPhoto" bson:"content_photo,omitempty"`
	ContentFile  File            `json:"contentFile" bson:"content_file,omitempty"`
	Likes        []bson.ObjectId `json:"likes" bson:"likes,omitempty"`
	Agreements   []bson.ObjectId `json:"agreements" bson:"agreements,omitempty"`
	Followers    []bson.ObjectId `json:"followers" bson:"followers,omitempty"`
	Comments     []Comment       `json:"comments" bson:"comments,omitempty"`
	Shares       []Share         `json:"shares" bson:"shares,omitempty"`
}

// Image ...
type Image struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name   string        `json:"name" bson:"name,omitempty"`
	Type   string        `json:"type" bson:"type,omitempty"`
	Size   int64         `json:"size" bson:"size,omitempty"`
	Width  int           `json:"width" bson:"width,omitempty"`
	Height int           `json:"height" bson:"height,omitempty"`
}

// File ...
type File struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name   string        `json:"name" bson:"name,omitempty"`
	Type   string        `json:"type" bson:"type,omitempty"`
	Size   int64         `json:"size" bson:"size,omitempty"`
	Width  int           `json:"width" bson:"width,omitempty"`
	Height int           `json:"height" bson:"height,omitempty"`
}

// Comment ...
type Comment struct {
	// ID     bson.ObjectId `json:"id" bson:"_id"`
	UserID bson.ObjectId `json:"userid" bson:"userid,omitempty"`
	Text   string        `json:"text" bson:"text,omitempty"`
}

// Share ...
type Share struct {
	Party string `json:"party" bson:"text"`
}

// NewPost creates a new User!
func NewPost(author, topic, category, text string) *Post {
	return &Post{
		Author:      bson.ObjectIdHex(author),
		Topic:       topic,
		Category:    category,
		ContentText: text,
		Date:        now(),
	}
}
