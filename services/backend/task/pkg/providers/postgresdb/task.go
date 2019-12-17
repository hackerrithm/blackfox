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
	"errors"

	cfg "github.com/hackerrithm/blackfox/services/backend/task/configs"
	"github.com/hackerrithm/blackfox/services/backend/task/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/task/pkg/engine"
	gorm "github.com/jinzhu/gorm"
	"golang.org/x/net/context"
)

type (
	taskRepository struct {
		db *gorm.DB
	}
)

const (
	taskCollection = "task"
)

var (
	config cfg.Config
)

func newTaskRepository(db *gorm.DB) engine.TaskRepository {
	return &taskRepository{db}
}

func (r taskRepository) Insert(c context.Context, p domain.Task) error {
	var err error
	err = r.db.Debug().Create(&p).Error
	if err != nil {
		return err
	}

	return nil
}

func (r taskRepository) Update(c context.Context, p domain.Task, id uint32) error {
	var t domain.Task
	r.db = r.db.Debug().Model(&domain.Task{}).Where("id = ?", id).Take(&domain.Task{}).UpdateColumns(
		map[string]interface{}{
			"text": p.Text,
		},
	)
	if r.db.Error != nil {
		return r.db.Error
	}
	// This is the display the updated user
	err := r.db.Debug().Model(&domain.Task{}).Where("id = ?", id).Take(&t).Error
	if err != nil {
		return err
	}
	return nil
}

func (r taskRepository) Query(c context.Context, query *engine.Query) []*domain.Task {
	// var err error
	// tasks := []domain.Task{}
	// err = r.db.Debug().Model(&User{}).Limit(100).Find(&tasks).Error
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (r taskRepository) FindOne(c context.Context, id uint32) (*domain.Task, error) {
	var err error
	var t *domain.Task
	err = r.db.Debug().Model(domain.Task{}).Where("id = ?", id).Take(&(t)).Error
	if err != nil {
		return &domain.Task{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &domain.Task{}, errors.New("User Not Found")
	}
	return t, err
}

// ListAllTasks used for finding all user tasks
// by the passed skip and take parameters
func (r taskRepository) ListAllTasks(ctx context.Context, skip uint64, take uint64) ([]domain.Task, error) {
	var err error
	tasks := []domain.Task{}
	err = r.db.Debug().Model(&domain.Task{}).Limit(100).Find(&tasks).Error
	if err != nil {
		return []domain.Task{}, err
	}
	return tasks, nil
}

func (r taskRepository) Remove(c context.Context, id uint32) (int64, error) {
	r.db = r.db.Debug().Model(&domain.Task{}).Where("id = ?", id).Take(&domain.Task{}).Delete(&domain.Task{})

	if r.db.Error != nil {
		return 0, r.db.Error
	}
	return r.db.RowsAffected, nil
}
