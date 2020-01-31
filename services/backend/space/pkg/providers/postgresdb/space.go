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
	"log"

	"golang.org/x/net/context"

	cfg "github.com/hackerrithm/blackfox/services/backend/space/configs"
	"github.com/hackerrithm/blackfox/services/backend/space/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/space/pkg/engine"
)

type (
	spaceRepository struct {
		db *sql.DB
	}
)

const (
	spaceCollection = "space"
)

var (
	config cfg.Config
)

func newSpaceRepository(db *sql.DB) engine.SpaceRepository {
	return &spaceRepository{db}
}

func (r spaceRepository) Insert(c context.Context, p domain.Space) error {
	log.Println("got in storage")

	// err = r.db.Debug().Create(&p).Error
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (r spaceRepository) Update(c context.Context, p domain.Space, id uint64) error {
	// var t domain.Space
	// r.db = r.db.Debug().Model(&domain.Space{}).Where("id = ?", id).Take(&domain.Space{}).UpdateColumns(
	// 	map[string]interface{}{
	// 		"creator":     p.Creator,
	// 		"description": p.Description,
	// 		"details":     p.Details,
	// 		"date":        p.Date,
	// 		"followers":   p.Followers,
	// 		"tags":        p.Tags,
	// 		"type":        p.Type,
	// 		"topic":       p.Topic,
	// 	},
	// )
	// if r.db.Error != nil {
	// 	return r.db.Error
	// }
	// // This is the display the updated user
	// err := r.db.Debug().Model(&domain.Space{}).Where("id = ?", id).Take(&t).Error
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (r spaceRepository) Query(c context.Context, query *engine.Query) []*domain.Space {
	return nil
}

func (r spaceRepository) FindOne(c context.Context, id uint64) (*domain.Space, error) {
	// var err error
	// var t *domain.Space
	// err = r.db.Debug().Model(domain.Space{}).Where("id = ?", id).Take(&(t)).Error
	// if err != nil {
	// 	return &domain.Space{}, err
	// }
	// if gorm.IsRecordNotFoundError(err) {
	// 	return &domain.Space{}, errors.New("User Not Found")
	// }
	return nil, nil
}

// ListAllSpaces used for finding all user spaces
// by the passed skip and take parameters
func (r spaceRepository) ListAllSpaces(ctx context.Context, skip uint64, take uint64) ([]domain.Space, error) {
	// var err error
	// tasks := []domain.Space{}
	// err = r.db.Debug().Model(&domain.Space{}).Limit(100).Find(&tasks).Error
	// if err != nil {
	// 	return []domain.Space{}, err
	// }
	return nil, nil
}

func (r spaceRepository) Remove(c context.Context, id uint64) (string, error) {
	// r.db = r.db.Debug().Model(&domain.Space{}).Where("id = ?", id).Take(&domain.Space{}).Delete(&domain.Space{})

	// if r.db.Error != nil {
	// 	return "", r.db.Error
	// }
	// return string(r.db.RowsAffected), nil
	return "", nil
}
