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

package postgresdb

// import (
// 	"fmt"
// 	"log"
// 	"sync"
//
// 

// 	gorm "github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/postgres" // _ postgres needed

// 	"github.com/hackerrithm/blackfox/services/backend/task/pkg/domain"
// 	"github.com/hackerrithm/blackfox/services/backend/task/pkg/engine"
// )

// type (
// 	storageFactory struct {
// 		session *gorm.DB
// 	}
// )

// var (
// 	taskRepositoryInstance engine.TaskRepository
// 	taskRepositoryOnce     sync.Once
// )

// // NewStorage creates a new instance of this mongodb storage factory
// func NewStorage(url ...string) (engine.StorageFactory, error) {
// 	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", url[0], url[1], url[2], url[3], url[4])
// 	db, err := gorm.Open("postgres", DBURL)
// 	if err != nil {
// 		fmt.Printf("Cannot connect to %s database", "postgres")
// 		log.Fatal("This is the error:", err)
// 	} else {
// 		fmt.Printf("We are connected to the %s database", "postgres")
// 	}
// 	return &storageFactory{db}, nil
// }

// // Automigrate ...
// func (s *storageFactory) Automigrate() {
// 	s.session.Debug().AutoMigrate(&domain.Task{})
// }

// // NewTaskRepository creates a new datastore Task repository
// func (s *storageFactory) NewTaskRepository() engine.TaskRepository {
// 	taskRepositoryOnce.Do(func() {
// 		taskRepositoryInstance = newTaskRepository(s.session)
// 	})
// 	return taskRepositoryInstance
// }

// func (s *storageFactory) Close() {
// 	s.session.Close()
// }
