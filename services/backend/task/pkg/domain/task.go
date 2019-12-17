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

package domain

import (
	"context"
	"time"
)

var (
	taskContextKey contextKey = "task"
)

type (
	contextKey string

	// Task is a task object
	// that has datatypes for
	// the task domain (Task)
	Task struct {
		ID   uint32    `gorm:"primary_key;auto_increment" json:"id"`
		Text string    `gorm:"size:255;not null;" json:"text"`
		Date time.Time `json:"date" bson:"date,omitempty"`
	}
)

// NewContext ...
func (u *Task) NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, taskContextKey, u)
}

// TaskFromContext gets task from context
func TaskFromContext(ctx context.Context) (*Task, bool) {
	u, ok := ctx.Value(taskContextKey).(*Task)
	return u, ok
}

// TaskMustFromContext gets task from context. if can't make panic
func TaskMustFromContext(ctx context.Context) *Task {
	u, ok := ctx.Value(taskContextKey).(*Task)
	if !ok {
		panic("task can't get from request's context")
	}
	return u
}

// NewTask creates a new Task!
func NewTask(text string) *Task {
	return &Task{
		Text: text,
		Date: now(),
	}
}
