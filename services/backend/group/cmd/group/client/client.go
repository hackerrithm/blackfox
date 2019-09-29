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

	"github.com/hackerrithm/blackfox/services/backend/group/pkg/domain"
	pb "github.com/hackerrithm/blackfox/services/backend/group/pkg/model"
)

// Client ...
type Client struct {
	conn    *grpc.ClientConn
	service pb.GroupServiceClient
}

// NewClient ...
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second*100))
	if err != nil {
		return nil, err
	}
	c := pb.NewGroupServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close ...
func (c *Client) Close() {
	c.conn.Close()
}

// Post ...
func (c *Client) Post(ctx context.Context, title, details, description, typ string, people []string) (string, error) {
	r, err := c.service.PostGroup(
		ctx,
		&pb.PostGroupRequest{
			Title:       title,
			Details:     details,
			Description: description,
			Type:        typ,
			People:      people,
		},
	)
	if err != nil {
		return "", err
	}

	log.Println(r.Group)
	return r.Group, nil
}

// Put ...
func (c *Client) Put(ctx context.Context, id, title, details, description, typ string, people []string) (string, error) {
	r, err := c.service.PutGroup(
		ctx,
		&pb.PutGroupRequest{
			Title:       title,
			Details:     details,
			Description: description,
			Type:        typ,
			People:      people,
		},
	)
	if err != nil {
		return "", err
	}
	return r.Group, nil
}

// Get ...
func (c *Client) Get(ctx context.Context, id string, userID uint64) (*domain.Group, error) {
	r, err := c.service.GetGroup(
		ctx,
		&pb.GetGroupRequest{Id: id, UserID: userID},
	)
	if err != nil {
		return nil, err
	}

	var people []bson.ObjectId
	for _, s := range r.Group.People {
		people = append(people, bson.ObjectIdHex(s))
	}

	return &domain.Group{
		ID:          bson.ObjectIdHex(r.Group.Id),
		Title:       r.Group.Title,
		Details:     r.Group.Details,
		Description: r.Group.Description,
		Type:        r.Group.Type,
		People:      people,
	}, nil
}

// GetMultiple is used to get the list of specified groups
func (c *Client) GetMultiple(ctx context.Context, skip uint64, take uint64) ([]domain.Group, error) {
	var people []bson.ObjectId

	r, err := c.service.GetMultipleGroups(
		ctx,
		&pb.GetMultipleGroupsRequest{
			Skip: skip,
			Take: take,
		},
	)
	if err != nil {
		return nil, err
	}

	groups := []domain.Group{}
	for _, a := range r.Groups {
		for i, s := range r.Groups {
			people = append(people, bson.ObjectIdHex(s.People[i]))
		}

		groups = append(groups, domain.Group{
			ID:          bson.ObjectIdHex(a.Id),
			Title:       a.Title,
			Details:     a.Details,
			Description: a.Description,
			Type:        a.Type,
			People:      people,
		})
	}
	return groups, nil
}

// Delete removes a group with passed identifier
func (c *Client) Delete(ctx context.Context, id string) (string, error) {
	r, err := c.service.DeleteGroup(
		ctx,
		&pb.DeleteGroupRequest{Id: id},
	)
	if err != nil {
		return "", err
	}
	return r.Id, nil
}
