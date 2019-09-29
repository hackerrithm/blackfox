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

package event

import (
	"github.com/hackerrithm/blackfox/services/backend/reaction/pkg/domain"
)

// ReactionEventStore ...
type ReactionEventStore interface {
	Close()
	PublishReactionCreated(reaction domain.Reaction) error
	SubscribeReactionCreated() (<-chan ReactionCreatedMessage, error)
	OnReactionCreated(f func(ReactionCreatedMessage)) error
}

var impl ReactionEventStore

// SetReactionEventStore ...
func SetReactionEventStore(es ReactionEventStore) {
	impl = es
}

// Close ...
func Close() {
	impl.Close()
}

// PublishReactionCreated ...
func PublishReactionCreated(reaction domain.Reaction) error {
	return impl.PublishReactionCreated(reaction)
}

// SubscribeReactionCreated ...
func SubscribeReactionCreated() (<-chan ReactionCreatedMessage, error) {
	return impl.SubscribeReactionCreated()
}

// OnReactionCreated ...
func OnReactionCreated(f func(ReactionCreatedMessage)) error {
	return impl.OnReactionCreated(f)
}
