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

package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/hackerrithm/blackfox/services/backend/task/pkg/domain"
	pb "github.com/hackerrithm/blackfox/services/backend/task/pkg/model"
)

// Client ...
type Client struct {
	conn    *grpc.ClientConn
	service pb.TaskServiceClient
}

// NewClient ...
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second*50))
	if err != nil {
		return nil, err
	}
	c := pb.NewTaskServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close ...
func (c *Client) Close() {
	c.conn.Close()
}

// Post ...
func (c *Client) Post(ctx context.Context, text string) (string, error) {
	fmt.Println("this is text: ", text)
	r, err := c.service.PostTask(
		ctx,
		&pb.PostTaskRequest{
			Text: text,
		},
	)
	if err != nil {
		return "this is not a insert error", err
	}

	log.Println(r.Task)
	return r.Task, nil
}

// Put ...
func (c *Client) Put(ctx context.Context, id uint32, text string) (string, error) {
	r, err := c.service.PutTask(
		ctx,
		&pb.PutTaskRequest{
			Id:   id,
			Text: text,
		},
	)
	if err != nil {
		return "", err
	}
	return r.Task, nil
}

// Get ...
func (c *Client) Get(ctx context.Context, id uint32, userID uint64) (*domain.Task, error) {
	r, err := c.service.GetTask(
		ctx,
		&pb.GetTaskRequest{Id: id, UserID: userID},
	)
	if err != nil {
		return nil, err
	}
	return &domain.Task{
		ID:   r.Task.Id,
		Text: r.Task.Text,
	}, nil
}

// GetMultiple is used to get the list of specified tasks
func (c *Client) GetMultiple(ctx context.Context, skip uint64, take uint64) ([]domain.Task, error) {
	r, err := c.service.GetMultipleTask(
		ctx,
		&pb.GetMultipleTaskRequest{
			Skip: skip,
			Take: take,
		},
	)
	if err != nil {
		return nil, err
	}

	tasks := []domain.Task{}
	for _, a := range r.Tasks {
		tasks = append(tasks, domain.Task{
			ID:   a.Id,
			Text: a.Text,
		})
	}
	return tasks, nil
}

// Delete removes a task with passed identifier
func (c *Client) Delete(ctx context.Context, id uint32) (uint32, error) {
	r, err := c.service.DeleteTask(
		ctx,
		&pb.DeleteTaskRequest{Id: id},
	)
	if err != nil {
		return 0, err
	}
	return r.Id, nil
}
