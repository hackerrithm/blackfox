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
	"gopkg.in/mgo.v2/bson"

	"github.com/hackerrithm/blackfox/services/backend/chat/pkg/domain"
	pb "github.com/hackerrithm/blackfox/services/backend/chat/pkg/model"
)

// Client ...
type Client struct {
	conn    *grpc.ClientConn
	service pb.ChatServiceClient
}

// NewClient ...
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second*100))
	if err != nil {
		return nil, err
	}
	c := pb.NewChatServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close ...
func (c *Client) Close() {
	c.conn.Close()
}

// Post ...
func (c *Client) Post(ctx context.Context, creator, topic, details, description, typ string, managers, followers, tags []string) (string, error) {
	r, err := c.service.Post(
		ctx,
		&pb.PostRequest{
			Creator:     creator,
			Topic:       topic,
			Details:     details,
			Description: description,
			Type:        typ,
			Followers:   followers,
			Managers:    managers,
			Tags:        tags,
		},
	)
	if err != nil {
		return "", err
	}

	log.Println(r.Chat)
	return r.Chat, nil
}

// Put ...
func (c *Client) Put(ctx context.Context, id, creator, topic, details, description, typ string, managers, followers, tags []string) (string, error) {
	r, err := c.service.Put(
		ctx,
		&pb.PutRequest{
			Id:          id,
			Creator:     creator,
			Topic:       topic,
			Details:     details,
			Description: description,
			Type:        typ,
			Followers:   followers,
			Managers:    managers,
			Tags:        tags,
		},
	)
	if err != nil {
		return "", err
	}
	return r.Chat, nil
}

// Get ...
func (c *Client) Get(ctx context.Context, id string, userID uint64) (*domain.Chat, error) {
	r, err := c.service.Get(
		ctx,
		&pb.GetRequest{Id: id, UserID: userID},
	)
	if err != nil {
		return nil, err
	}

	var mngs, fllwrs []bson.ObjectId
	for _, m := range r.Chat.Managers {
		mngs = append(mngs, bson.ObjectIdHex(m))
	}
	for _, f := range r.Chat.Followers {
		fllwrs = append(fllwrs, bson.ObjectIdHex(f))
	}

	return &domain.Chat{
		ID:          bson.ObjectIdHex(r.Chat.Id),
		Creator:     bson.ObjectIdHex(r.Chat.Creator),
		Topic:       r.Chat.Topic,
		Details:     r.Chat.Details,
		Description: r.Chat.Description,
		Type:        r.Chat.Type,
		Followers:   fllwrs,
		Managers:    mngs,
		Tags:        r.Chat.Tags,
	}, nil
}

// GetMultiple is used to get the list of specified chats
func (c *Client) GetMultiple(ctx context.Context, skip uint64, take uint64) ([]domain.Chat, error) {
	var mngs, fllwrs []bson.ObjectId

	r, err := c.service.GetMultiple(
		ctx,
		&pb.GetMultipleRequest{
			Skip: skip,
			Take: take,
		},
	)
	if err != nil {
		return nil, err
	}

	chats := []domain.Chat{}
	for _, a := range r.Chats {
		for _, m := range a.Managers {
			mngs = append(mngs, bson.ObjectIdHex(m))
		}
		for _, f := range a.Followers {
			fllwrs = append(fllwrs, bson.ObjectIdHex(f))
		}

		chats = append(chats, domain.Chat{
			ID:          bson.ObjectIdHex(a.Id),
			Creator:     bson.ObjectIdHex(a.Creator),
			Topic:       a.Topic,
			Details:     a.Details,
			Description: a.Description,
			Type:        a.Type,
			Followers:   fllwrs,
			Managers:    mngs,
			Tags:        a.Tags,
		})
	}
	return chats, nil
}

// Delete removes a chat with passed identifier
func (c *Client) Delete(ctx context.Context, id string) (string, error) {
	r, err := c.service.Delete(
		ctx,
		&pb.DeleteRequest{Id: id},
	)
	if err != nil {
		return "", err
	}
	return r.Id, nil
}
