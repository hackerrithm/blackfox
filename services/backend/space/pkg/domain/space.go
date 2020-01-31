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
)

var spaceContextKey contextKey = "space"

type contextKey string

// Space ...
// type Space struct {
// 	ID          bson.ObjectId   `json:"id" bson:"_id,omitempty"`
// 	Creator     bson.ObjectId   `json:"author" bson:"author,omitempty"`
// 	Managers    []bson.ObjectId `json:"managers" bson:"managers,omitempty"`
// 	Topic       string          `json:"topic" bson:"topic,omitempty"`
// 	Details     string          `json:"details" bson:"details,omitempty"`
// 	Description string          `json:"description" bson:"description,omitempty"`
// 	Type        string          `json:"type" bson:"type,omitempty"`
// 	Tags        []string        `json:"tags" bson:"tags,omitempty"`
// 	Date        time.Time       `json:"timestamp" bson:"timestamp,omitempty"`
// 	Followers   []bson.ObjectId `json:"followers" bson:"followers,omitempty"`
// }

// Space ...
// type Space struct {
// 	ID          uint64         `gorm:"primary_key;auto_increment" json:"id"`
// 	Creator     string         `gorm:"type:varchar(100)" json:"creator"`
// 	Managers    pq.StringArray `gorm:"type:varchar(64)[]"`
// 	Topic       string         `json:"topic" gorm:"type:varchar(50)"`
// 	Details     string         `json:"details" gorm:"type:varchar(100)"`
// 	Description string         `json:"description" gorm:"type:varchar(100)"`
// 	Type        string         `json:"type" gorm:"type:varchar(45)"`
// 	Tags        pq.StringArray `gorm:"type:varchar(64)[]"`
// 	Date        time.Time      `json:"timestamp"`
// 	Followers   pq.StringArray `gorm:"type:varchar(64)[]"`
// }

// Space ...
type Space struct {
	ID          uint64    `json:"id"`
	Creator     string    `json:"creator"`
	Managers    []string  `json:"managers"`
	Topic       string    `json:"topic"`
	Details     string    `json:"details"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Tags        []string  `json:"tags"`
	Date        time.Time `json:"timestamp"`
	Followers   []string  `json:"followers"`
}

// NewSpace creates a new Space!
// func NewSpace(creator, topic, details, description, typ string, managers, followers, tags []string) *Space {
// 	var mngs, fllwrs []bson.ObjectId
// 	for _, m := range managers {
// 		mngs = append(mngs, bson.ObjectIdHex(m))
// 	}
// 	for _, f := range followers {
// 		fllwrs = append(fllwrs, bson.ObjectIdHex(f))
// 	}
// 	return &Space{
// 		Creator:     bson.ObjectIdHex(creator),
// 		Topic:       topic,
// 		Details:     details,
// 		Description: description,
// 		Type:        typ,
// 		Followers:   fllwrs,
// 		Managers:    mngs,
// 		Tags:        tags,
// 		Date:        now(),
// 	}
// }

// NewSpace ...
func NewSpace(creator, topic, details, description, typ string, managers, followers, tags []string) *Space {
	return &Space{
		Creator:     creator,
		Topic:       topic,
		Details:     details,
		Description: description,
		Type:        typ,
		Followers:   followers,
		Managers:    managers,
		Tags:        tags,
		Date:        now(),
	}
}
