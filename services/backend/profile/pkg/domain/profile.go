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
	"context"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var profileContextKey contextKey = "profile"

type (
	contextKey string

	// Profile ...
	Profile struct {
		ID              bson.ObjectId   `json:"id" bson:"_id,omitempty"`
		UserName        string          `json:"username" bson:"username,omitempty"`
		Level           string          `json:"level" bson:"level,omitempty"`
		Rings           int32           `json:"rings" bson:"rings,omitempty"`
		About           string          `json:"about" bson:"about,omitempty"`
		ProfileImage    Image           `json:"profileImage" bson:"profile_image,omitempty"`
		BackgroundImage Image           `json:"backgroundImage" bson:"background_image,omitempty"`
		Followers       []bson.ObjectId `json:"followers" bson:"followers,omitempty"`
		Following       []bson.ObjectId `json:"following" bson:"following,omitempty"`
		DateLastUpdated time.Time       `json:"dateLastUpdated" bson:"date_last_updated,omitempty"`
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

// NewContext ...
func (u *Profile) NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, profileContextKey, u)
}

// ProfileFromContext gets profile from context
func ProfileFromContext(ctx context.Context) (*Profile, bool) {
	u, ok := ctx.Value(profileContextKey).(*Profile)
	return u, ok
}

// ProfileMustFromContext gets profile from context. if can't make panic
func ProfileMustFromContext(ctx context.Context) *Profile {
	u, ok := ctx.Value(profileContextKey).(*Profile)
	if !ok {
		panic("profile can't get from request's context")
	}
	return u
}

// NewProfile creates a new Profile!
func NewProfile(username, level, about string, followers []string, following []string, rings int32) *Profile {
	var fllwrs []bson.ObjectId
	var fllwing []bson.ObjectId

	for _, f := range followers {
		fllwrs = append(fllwrs, bson.ObjectIdHex(f))
	}

	for _, f := range following {
		fllwing = append(fllwing, bson.ObjectIdHex(f))
	}
	return &Profile{
		UserName:        username,
		About:           about,
		Rings:           rings,
		Followers:       fllwrs,
		Following:       fllwing,
		DateLastUpdated: now(),
	}
}
