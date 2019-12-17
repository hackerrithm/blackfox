// // Copyright 2019 kemar
// //
// // Licensed under the Apache License, Version 2.0 (the "License");
// // you may not use this file except in compliance with the License.
// // You may obtain a copy of the License at
// //
// //     http://www.apache.org/licenses/LICENSE-2.0
// //
// // Unless required by applicable law or agreed to in writing, software
// // distributed under the License is distributed on an "AS IS" BASIS,
// // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// // See the License for the specific language governing permissions and
// // limitations under the License.

// package mongodb

// import (
// 	"log"

// 	"golang.org/x/net/context"
// 	mgo "gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"

// 	cfg "github.com/hackerrithm/blackfox/services/backend/task/configs"
// 	"github.com/hackerrithm/blackfox/services/backend/task/pkg/domain"
// 	"github.com/hackerrithm/blackfox/services/backend/task/pkg/engine"
// )

// type (
// 	taskRepository struct {
// 		session *mgo.Session
// 	}
// )

// const (
// 	taskCollection = "task"
// )

// var (
// 	config cfg.Config
// )

// func newTaskRepository(session *mgo.Session) engine.TaskRepository {
// 	return &taskRepository{session}
// }

// func (r taskRepository) Insert(c context.Context, p domain.Task) error {
// 	s := r.session.Clone()
// 	defer s.Close()

// 	var task domain.Task

// 	task.Text = p.Text
// 	task.Date = p.Date

// 	col := s.DB(config.MongoDB).C(config.MongoCollection)
// 	err := col.Insert(&task)
// 	if err != nil {
// 		log.Println("insert error ", err)
// 	}

// 	return nil
// }

// func (r taskRepository) Update(c context.Context, p domain.Task, id string) error {
// 	s := r.session.Clone()
// 	defer s.Close()

// 	var task domain.Task

// 	task.Text = p.Text
// 	task.Date = p.Date

// 	col := s.DB(config.MongoDB).C(config.MongoCollection)
// 	err := col.Update(bson.M{"_id": toBSON(id)}, p)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r taskRepository) Query(c context.Context, query *engine.Query) []*domain.Task {
// 	s := r.session.Clone()
// 	defer s.Close()

// 	col := s.DB(config.MongoDB).C(config.MongoCollection)
// 	p := []*domain.Task{}
// 	q := translateQuery(col, query)
// 	q.All(&p)

// 	return nil
// }

// func (r taskRepository) FindOne(c context.Context, id string) (*domain.Task, error) {
// 	s := r.session.Clone()
// 	defer s.Close()

// 	var task *domain.Task
// 	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"_id": toBSON(id)}).One(&task)

// 	if err != nil {
// 		return nil, nil
// 	}

// 	return task, nil
// }

// // ListAllTasks used for finding all user tasks
// // by the passed skip and take parameters
// func (r taskRepository) ListAllTasks(ctx context.Context, skip uint64, take uint64) ([]domain.Task, error) {
// 	s := r.session.Clone()
// 	defer s.Close()
// 	log.Println("got in repo")

// 	var tasks []domain.Task

// 	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(nil).All(&tasks)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return tasks, nil
// }

// func (r taskRepository) Remove(c context.Context, id string) (string, error) {
// 	s := r.session.Clone()
// 	defer s.Close()

// 	err := s.DB(config.MongoDB).C(config.MongoCollection).Remove(bson.M{"_id": toBSON(id)}) //.One(&task)

// 	if err != nil {
// 		return "", nil
// 	}

// 	return "", nil
// }
