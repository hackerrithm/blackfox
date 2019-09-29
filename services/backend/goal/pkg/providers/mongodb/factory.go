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
	"sync"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/hackerrithm/blackfox/services/backend/goal/pkg/engine"
)

type (
	storageFactory struct {
		session *mgo.Session
	}
)

var (
	goalRepositoryInstance engine.GoalRepository
	goalRepositoryOnce     sync.Once
)

// NewStorage creates a new instance of this mongodb storage factory
func NewStorage(url ...string) (engine.StorageFactory, error) {
	info := &mgo.DialInfo{
		Addrs:    []string{url[0]},
		Timeout:  60 * time.Second,
		Database: url[1],
		Username: url[2],
		Password: url[3],
	}
	config.MongoDB = url[1]
	config.MongoCollection = url[4]

	log.Println(info)

	session, err := mgo.DialWithInfo(info)
	// session, err := mgo.Dial(url[5]) //mgo.DialWithInfo(info)

	if err != nil {
		return nil, err
	}
	log.Println("connected")

	ensureIndexes(session)
	return &storageFactory{session}, nil
}

// NewGoalRepository creates a new datastore Goal repository
func (f *storageFactory) NewGoalRepository() engine.GoalRepository {
	goalRepositoryOnce.Do(func() {
		goalRepositoryInstance = newGoalRepository(f.session)
	})
	return goalRepositoryInstance
}

func ensureIndexes(s *mgo.Session) {
	index := mgo.Index{
		Key:        []string{"date"},
		Background: true,
	}
	c := s.DB("test1").C(goalCollection)
	c.EnsureIndex(index)
}

func (f *storageFactory) Close() {
	f.session.Close()
}
