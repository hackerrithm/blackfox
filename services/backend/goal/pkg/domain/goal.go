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

var goalContextKey contextKey = "goal"

type contextKey string

// Goal ...
type Goal struct {
	ID           bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Creator      bson.ObjectId   `json:"creator" bson:"creator,omitempty"`
	Participants []bson.ObjectId `json:"participants" bson:"participants,omitempty"`
	Likes        []bson.ObjectId `json:"likes" bson:"likes,omitempty"`
	Watchers     []bson.ObjectId `json:"watchers" bson:"watchers,omitempty"`
	Aim          string          `json:"aim" bson:"aim,omitempty"`
	Reason       string          `json:"reason" bson:"reason"`
	Details      string          `json:"details" bson:"details,omitempty"`
	Inspiration  string          `json:"inspiration" bson:"inspiration,omitempty"`
	Type         string          `json:"type" bson:"type,omitempty"`
	Tags         []string        `json:"tags" bson:"tags,omitempty"`
	SimilarGoals []string        `json:"similarGoals" bson:"similarGoals,omitempty"`
	Journey      Journey         `json:"journey" bson:"journey"`
	IsAchieved   bool            `json:"isAchieved" bson:"isAchieved"`
	IsPrivate    bool            `json:"isPrivate" bson:"isPrivate,omitempty"`
	Date         time.Time       `json:"timestamp" bson:"timestamp,omitempty"`
}

// Journey ...
type Journey struct {
	IsComplete   bool      `json:"isComplete" bson:"isComplete,omitempty"`
	IsInProgress bool      `json:"isInProgress" bson:"isInProgress,omitempty"`
	IsStarted    bool      `json:"isStarted" bson:"isStarted,omitempty"`
	Details      string    `json:"details" bson:"details,omitempty"`
	Type         string    `json:"type" bson:"type,omitempty"`
	Steps        []string  `json:"steps" bson:"steps,omitempty"`
	Progress     int32     `json:"progress" bson:"progress,omitempty"`
	StartDate    time.Time `json:"startdate" bson:"startdate,omitempty"`
	DueDate      time.Time `json:"duedate" bson:"duedate,omitempty"`
}

// NewGoal creates a new Goal!
func NewGoal(creator, aim, reason, details, typ string, tags []string) *Goal {
	return &Goal{
		Creator: bson.ObjectIdHex(creator),
		Aim:     aim,
		Reason:  reason,
		Details: details,
		Type:    typ,
		Tags:    tags,
		Journey: Journey{},
		Date:    now(),
	}
}
