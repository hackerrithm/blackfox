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

package mongodb

import (
	"log"

	"golang.org/x/net/context"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	cfg "github.com/hackerrithm/blackfox/services/backend/chat/configs"
	"github.com/hackerrithm/blackfox/services/backend/chat/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/chat/pkg/engine"
)

type (
	chatRepository struct {
		session *mgo.Session
	}
)

const (
	chatCollection = "chat"
)

var (
	config cfg.Config
)

func newChatRepository(session *mgo.Session) engine.ChatRepository {
	return &chatRepository{session}
}

func (r chatRepository) Insert(c context.Context, p domain.Chat) error {
	s := r.session.Clone()
	defer s.Close()

	var chat domain.Chat

	chat.Creator = p.Creator
	chat.Description = p.Description
	chat.Details = p.Details
	chat.Managers = p.Managers
	chat.Date = p.Date
	chat.Followers = p.Followers
	chat.Tags = p.Tags
	chat.Type = p.Type
	chat.Topic = p.Topic
	chat.Date = p.Date

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Insert(&chat)
	if err != nil {
		log.Println("insert error ", err)
	}

	return nil
}

func (r chatRepository) Update(c context.Context, p domain.Chat, id string) error {
	s := r.session.Clone()
	defer s.Close()

	var chat domain.Chat

	chat.Creator = p.Creator
	chat.Description = p.Description
	chat.Details = p.Details
	chat.Managers = p.Managers
	chat.Date = p.Date
	chat.Followers = p.Followers
	chat.Tags = p.Tags
	chat.Type = p.Type
	chat.Topic = p.Topic
	chat.Date = p.Date

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Update(bson.M{"_id": toBSON(id)}, p)
	if err != nil {
		return err
	}

	return nil
}

func (r chatRepository) Query(c context.Context, query *engine.Query) []*domain.Chat {
	s := r.session.Clone()
	defer s.Close()

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	p := []*domain.Chat{}
	q := translateQuery(col, query)
	q.All(&p)

	return nil
}

func (r chatRepository) FindOne(c context.Context, id string) (*domain.Chat, error) {
	s := r.session.Clone()
	defer s.Close()

	var chat *domain.Chat
	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"_id": toBSON(id)}).One(&chat)

	if err != nil {
		return nil, nil
	}

	return chat, nil
}

// ListAllChats used for finding all user chats
// by the passed skip and take parameters
func (r chatRepository) ListAllChats(ctx context.Context, skip uint64, take uint64) ([]domain.Chat, error) {
	s := r.session.Clone()
	defer s.Close()
	log.Println("got in repo")

	var chats []domain.Chat

	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(nil).All(&chats)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (r chatRepository) Remove(c context.Context, id string) (string, error) {
	s := r.session.Clone()
	defer s.Close()

	err := s.DB(config.MongoDB).C(config.MongoCollection).Remove(bson.M{"_id": toBSON(id)}) //.One(&chat)

	if err != nil {
		return "", nil
	}

	return "", nil
}
