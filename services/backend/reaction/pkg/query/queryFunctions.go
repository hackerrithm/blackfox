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

package query

import (
	"context"
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/hackerrithm/blackfox/services/backend/reaction/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/reaction/pkg/event"
	"github.com/hackerrithm/blackfox/services/backend/reaction/pkg/search"
)

// OnReactionCreated ...
func OnReactionCreated(m event.ReactionCreatedMessage) {
	// Index reaction for searching
	reaction := domain.Reaction{
		PersonID:    bson.ObjectIdHex(m.PersonID),
		Details:     m.Details,
		Description: m.Description,
		Type:        m.Type,
		Date:        m.Date,
	}
	if err := search.InsertReaction(context.Background(), reaction); err != nil {
		log.Println(err)
	}
}
