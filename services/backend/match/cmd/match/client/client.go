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

	"github.com/hackerrithm/blackfox/services/backend/match/pkg/domain"
	pb "github.com/hackerrithm/blackfox/services/backend/match/pkg/model"
)

// Client ...
type Client struct {
	conn    *grpc.ClientConn
	service pb.MatchServiceClient
}

// NewClient ...
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second*100))
	if err != nil {
		return nil, err
	}
	c := pb.NewMatchServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close ...
func (c *Client) Close() {
	c.conn.Close()
}

// Post ...
func (c *Client) Post(ctx context.Context, personID, details, description, typ string, similarities []string) (string, error) {
	r, err := c.service.PostMatch(
		ctx,
		&pb.PostMatchRequest{
			PersonID:     personID,
			Details:      details,
			Description:  description,
			Type:         typ,
			Similarities: similarities,
		},
	)
	if err != nil {
		return "", err
	}

	log.Println(r.Match)
	return r.Match, nil
}

// Put ...
func (c *Client) Put(ctx context.Context, id, personID, details, description, typ string, similarities []string) (string, error) {
	r, err := c.service.PutMatch(
		ctx,
		&pb.PutMatchRequest{
			PersonID:     personID,
			Details:      details,
			Description:  description,
			Type:         typ,
			Similarities: similarities,
		},
	)
	if err != nil {
		return "", err
	}
	return r.Match, nil
}

// Get ...
func (c *Client) Get(ctx context.Context, id string, userID uint64) (*domain.Match, error) {
	r, err := c.service.GetMatch(
		ctx,
		&pb.GetMatchRequest{Id: id, UserID: userID},
	)
	if err != nil {
		return nil, err
	}

	var smlrts []string
	for _, s := range r.Match.Managers {
		smlrts = append(smlrts, s)
	}

	return &domain.Match{
		ID:           bson.ObjectIdHex(r.Match.Id),
		PersonID:     bson.ObjectIdHex(r.Match.PersonID),
		Details:      r.Match.Details,
		Description:  r.Match.Description,
		Type:         r.Match.Type,
		Similarities: smlrts,
	}, nil
}

// GetMultiple is used to get the list of specified matchs
func (c *Client) GetMultiple(ctx context.Context, similarities []string, skip uint64, take uint64) ([]domain.MatchedUser, error) {
	r, err := c.service.GetMultipleMatches(
		ctx,
		&pb.GetMultipleMatchesRequest{
			Similarities: similarities,
			Skip:         skip,
			Take:         take,
		},
	)
	if err != nil {
		return nil, err
	}

	matches := []domain.MatchedUser{}
	for _, a := range r.Matches {
		matches = append(matches, domain.MatchedUser{
			ID:           bson.ObjectIdHex(a.Id),
			Name:         a.Name,
			Username:     a.Username,
			Firstname:    a.Firstname,
			Lastname:     a.Lastname,
			Gender:       a.Gender,
			EmailAddress: a.EmailAddress,
			Type:         a.Type,
		})
	}
	return matches, nil
}

// Delete removes a match with passed identifier
func (c *Client) Delete(ctx context.Context, id string) (string, error) {
	r, err := c.service.DeleteMatch(
		ctx,
		&pb.DeleteMatchRequest{Id: id},
	)
	if err != nil {
		return "", err
	}
	return r.Id, nil
}
