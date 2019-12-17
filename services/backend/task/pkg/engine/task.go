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

package engine

import (
	"sync"

	"golang.org/x/net/context"

	"github.com/hackerrithm/blackfox/services/backend/task/pkg/domain"
)

type (
	// Task ...
	Task interface {
		Insert(ctx context.Context, text string) error

		// Update is the update-a-task use-case
		Update(ctx context.Context, id uint32, text string) error

		// Query is the list-the-tasks use-case
		Query(ctx context.Context, skip uint64, take uint64) ([]domain.Task, error)

		// FindOne ...
		FindOne(ctx context.Context, id uint32) (*domain.Task, error)

		// RemoveDelete ...
		Remove(ctx context.Context, id uint32) (int64, error)

		// ListAllTasks ...
		ListAllTasks(ctx context.Context, skip, take uint64) (*[]domain.Task, error)
	}

	task struct {
		repository TaskRepository
	}
)

var (
	taskInstance Task
	taskOnce     sync.Once
)

func (f *engineFactory) NewTask() Task {
	taskOnce.Do(func() {
		taskInstance = &task{
			repository: f.NewTaskRepository(),
			// jwt:        f.JWTSignParser,
		}
	})
	return taskInstance
}
