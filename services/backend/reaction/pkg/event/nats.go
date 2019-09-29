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
	"bytes"
	"encoding/gob"
	"time"

	"github.com/nats-io/go-nats"

	"github.com/hackerrithm/blackfox/services/backend/reaction/pkg/domain"
)

// NatsEventStore ...
type NatsEventStore struct {
	nc                          *nats.Conn
	reactionCreatedSubscription *nats.Subscription
	reactionCreatedChan         chan ReactionCreatedMessage
}

// NewNats ...
func NewNats(url string) (*NatsEventStore, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsEventStore{nc: nc}, nil
}

// SubscribeReactionCreated ...
func (e *NatsEventStore) SubscribeReactionCreated() (<-chan ReactionCreatedMessage, error) {
	m := ReactionCreatedMessage{}
	e.reactionCreatedChan = make(chan ReactionCreatedMessage, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	e.reactionCreatedSubscription, err = e.nc.ChanSubscribe(m.Key(), ch)
	if err != nil {
		return nil, err
	}
	// Decode message
	go func() {
		for {
			select {
			case msg := <-ch:
				e.readMessage(msg.Data, &m)
				e.reactionCreatedChan <- m
			}
		}
	}()
	return (<-chan ReactionCreatedMessage)(e.reactionCreatedChan), nil
}

// OnReactionCreated ...
func (e *NatsEventStore) OnReactionCreated(f func(ReactionCreatedMessage)) (err error) {
	m := ReactionCreatedMessage{}
	e.reactionCreatedSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		e.readMessage(msg.Data, &m)
		f(m)
	})
	return
}

// Close ...
func (e *NatsEventStore) Close() {
	if e.nc != nil {
		e.nc.Close()
	}
	if e.reactionCreatedSubscription != nil {
		e.reactionCreatedSubscription.Unsubscribe()
	}
	close(e.reactionCreatedChan)
}

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

// PublishReactionCreated ...
func (e *NatsEventStore) PublishReactionCreated(reaction domain.Reaction) error {
	formatedTime := reaction.Date.Format(time.RFC1123)
	x, _ := time.Parse(layoutISO, formatedTime)
	m := ReactionCreatedMessage{reaction.ID.Hex(), reaction.PersonID.Hex(), reaction.Description, reaction.Type, reaction.Details, x}
	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

// writeMessage ...maybe changed back to mq
func (e *NatsEventStore) writeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// readMessage ...
func (e *NatsEventStore) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}
