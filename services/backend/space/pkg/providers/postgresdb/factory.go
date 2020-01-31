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

package postgresdb

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq" // _ postgres needed

	"github.com/hackerrithm/blackfox/services/backend/space/pkg/engine"
)

type (
	storageFactory struct {
		db *sql.DB
	}
)

var (
	spaceRepositoryInstance engine.SpaceRepository
	spaceRepositoryOnce     sync.Once
)

// NewStorage creates a new instance of this postgres storage factory
func NewStorage(url ...string) (engine.StorageFactory, error) {
	// db, err := gorm.Open("postgres", "postgres://kemar:password@blackfox-postgres/postgres?sslmode=disable")
	// if err != nil {
	// 	return nil, err
	// }
	// defer db.Close()
	log.Println("url: ", url)
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", url[0], url[1], url[2], url[3], url[4])
	// DBURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", "blackfox-postgres", 5432, "kemar", "password", "blackfox_database_postgres")
	// DBURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", "postgres", 5432, "postgres", "postgres", "postgres")
	db, err := sql.Open("postgres", DBURL)
	if err != nil {
		log.Fatal("This is the error:", err)
	}
	// defer db.Close()
	return &storageFactory{db}, nil
}

// Automigrate ...
func (s *storageFactory) Automigrate() {
	//s.db.Debug().AutoMigrate(&domain.Space{})
}

// NewSpaceRepository creates a new datastore Space repository
func (s *storageFactory) NewSpaceRepository() engine.SpaceRepository {
	spaceRepositoryOnce.Do(func() {
		spaceRepositoryInstance = newSpaceRepository(s.db)
	})
	return spaceRepositoryInstance
}

// Close ...
func (s *storageFactory) Close() {
	s.db.Close()
}
