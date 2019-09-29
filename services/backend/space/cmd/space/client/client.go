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

	"github.com/hackerrithm/blackfox/services/backend/space/pkg/domain"
	pb "github.com/hackerrithm/blackfox/services/backend/space/pkg/model"
)

// Client ...
type Client struct {
	conn    *grpc.ClientConn
	service pb.SpaceServiceClient
}

// NewClient ...
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second*100))
	if err != nil {
		return nil, err
	}
	c := pb.NewSpaceServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close ...
func (c *Client) Close() {
	c.conn.Close()
}

// Post ...
func (c *Client) Post(ctx context.Context, creator, topic, details, description, typ string, managers, followers, tags []string) (string, error) {
	r, err := c.service.PostSpace(
		ctx,
		&pb.PostSpaceRequest{
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

	log.Println(r.Space)
	return r.Space, nil
}

// Put ...
func (c *Client) Put(ctx context.Context, id, creator, topic, details, description, typ string, managers, followers, tags []string) (string, error) {
	r, err := c.service.PutSpace(
		ctx,
		&pb.PutSpaceRequest{
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
	return r.Space, nil
}

// Get ...
func (c *Client) Get(ctx context.Context, id string, userID uint64) (*domain.Space, error) {
	r, err := c.service.GetSpace(
		ctx,
		&pb.GetSpaceRequest{Id: id, UserID: userID},
	)
	if err != nil {
		return nil, err
	}

	var mngs, fllwrs []bson.ObjectId
	for _, m := range r.Space.Managers {
		mngs = append(mngs, bson.ObjectIdHex(m))
	}
	for _, f := range r.Space.Followers {
		fllwrs = append(fllwrs, bson.ObjectIdHex(f))
	}

	return &domain.Space{
		ID:          bson.ObjectIdHex(r.Space.Id),
		Creator:     bson.ObjectIdHex(r.Space.Creator),
		Topic:       r.Space.Topic,
		Details:     r.Space.Details,
		Description: r.Space.Description,
		Type:        r.Space.Type,
		Followers:   fllwrs,
		Managers:    mngs,
		Tags:        r.Space.Tags,
	}, nil
}

// GetMultiple is used to get the list of specified spaces
func (c *Client) GetMultiple(ctx context.Context, skip uint64, take uint64) ([]domain.Space, error) {
	var mngs, fllwrs []bson.ObjectId

	r, err := c.service.GetMultipleSpaces(
		ctx,
		&pb.GetMultipleSpacesRequest{
			Skip: skip,
			Take: take,
		},
	)
	if err != nil {
		return nil, err
	}

	spaces := []domain.Space{}
	for _, a := range r.Spaces {
		for _, m := range a.Managers {
			mngs = append(mngs, bson.ObjectIdHex(m))
		}
		for _, f := range a.Followers {
			fllwrs = append(fllwrs, bson.ObjectIdHex(f))
		}

		spaces = append(spaces, domain.Space{
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
	return spaces, nil
}

// Delete removes a space with passed identifier
func (c *Client) Delete(ctx context.Context, id string) (string, error) {
	r, err := c.service.DeleteSpace(
		ctx,
		&pb.DeleteSpaceRequest{Id: id},
	)
	if err != nil {
		return "", err
	}
	return r.Id, nil
}
