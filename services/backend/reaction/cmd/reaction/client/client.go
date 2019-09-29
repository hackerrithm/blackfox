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

	"github.com/hackerrithm/blackfox/services/backend/reaction/pkg/domain"
	pb "github.com/hackerrithm/blackfox/services/backend/reaction/pkg/model"
)

// Client ...
type Client struct {
	conn    *grpc.ClientConn
	service pb.ReactionServiceClient
}

// NewClient ...
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second*100))
	if err != nil {
		return nil, err
	}
	c := pb.NewReactionServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close ...
func (c *Client) Close() {
	c.conn.Close()
}

// Post ...
func (c *Client) Post(ctx context.Context, personID, details, description, typ string) (string, error) {
	r, err := c.service.PostReaction(
		ctx,
		&pb.PostReactionRequest{
			PersonID:    personID,
			Details:     details,
			Description: description,
			Type:        typ,
		},
	)
	if err != nil {
		return "", err
	}

	log.Println(r.Reaction)
	return r.Reaction, nil
}

// Put ...
func (c *Client) Put(ctx context.Context, id, personID, details, description, typ string) (string, error) {
	r, err := c.service.PutReaction(
		ctx,
		&pb.PutReactionRequest{
			PersonID:    personID,
			Details:     details,
			Description: description,
			Type:        typ,
		},
	)
	if err != nil {
		return "", err
	}
	return r.Reaction, nil
}

// Get ...
func (c *Client) Get(ctx context.Context, id string, userID uint64) (*domain.Reaction, error) {
	r, err := c.service.GetReaction(
		ctx,
		&pb.GetReactionRequest{Id: id, UserID: userID},
	)
	if err != nil {
		return nil, err
	}

	return &domain.Reaction{
		ID:          bson.ObjectIdHex(r.Reaction.Id),
		PersonID:    bson.ObjectIdHex(r.Reaction.PersonID),
		Details:     r.Reaction.Details,
		Description: r.Reaction.Description,
		Type:        r.Reaction.Type,
	}, nil
}

// GetMultiple is used to get the list of specified reactions
func (c *Client) GetMultiple(ctx context.Context, skip uint64, take uint64) ([]domain.Reaction, error) {

	r, err := c.service.GetMultipleReactiones(
		ctx,
		&pb.GetMultipleReactionesRequest{
			Skip: skip,
			Take: take,
		},
	)
	if err != nil {
		return nil, err
	}

	reactiones := []domain.Reaction{}
	for _, a := range r.Reactiones {
		reactiones = append(reactiones, domain.Reaction{
			ID:          bson.ObjectIdHex(a.Id),
			PersonID:    bson.ObjectIdHex(a.PersonID),
			Details:     a.Details,
			Description: a.Description,
			Type:        a.Type,
		})
	}
	return reactiones, nil
}

// Delete removes a reaction with passed identifier
func (c *Client) Delete(ctx context.Context, id string) (string, error) {
	r, err := c.service.DeleteReaction(
		ctx,
		&pb.DeleteReactionRequest{Id: id},
	)
	if err != nil {
		return "", err
	}
	return r.Id, nil
}
