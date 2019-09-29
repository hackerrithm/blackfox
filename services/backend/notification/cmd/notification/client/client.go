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
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/hackerrithm/blackfox/services/backend/notification/pkg/domain"
	pb "github.com/hackerrithm/blackfox/services/backend/notification/pkg/model"
)

// Client ...
type Client struct {
	conn    *grpc.ClientConn
	service pb.NotificationServiceClient
}

// NewClient ...
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		return nil, err
	}
	c := pb.NewNotificationServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close ...
func (c *Client) Close() {
	c.conn.Close()
}

// SendNotification ...
func (c *Client) SendNotification(ctx context.Context, service, topic, body string) (*domain.Notification, error) {
	r, err := c.service.SendNotification(
		ctx,
		&pb.SendNotificationRequest{
			Service: service,
			Topic:   topic,
			Body:    body,
		},
	)
	if err != nil {
		return nil, err
	}

	log.Println(r.Notification)
	return &domain.Notification{
		Service: service,
		Topic:   topic,
		Body:    body,
	}, nil
}
