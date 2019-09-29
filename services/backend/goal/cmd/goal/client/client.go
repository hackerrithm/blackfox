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

	"github.com/hackerrithm/blackfox/services/backend/goal/pkg/domain"
	pb "github.com/hackerrithm/blackfox/services/backend/goal/pkg/model"
)

// Client ...
type Client struct {
	conn    *grpc.ClientConn
	service pb.GoalServiceClient
}

// NewClient ...
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		return nil, err
	}
	c := pb.NewGoalServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close ...
func (c *Client) Close() {
	c.conn.Close()
}

// Post ...
func (c *Client) Post(ctx context.Context, creator, aim, reason, details, typ, journeyType string, tags []string) (string, error) {
	var journey = pb.Journey{}

	journey.Typ = journeyType
	journey.IsInProgress = false
	journey.IsStarted = false
	journey.IsComplete = false
	journey.Details = ""
	journey.Steps = []string{}

	r, err := c.service.PostGoal(
		ctx,
		&pb.PostGoalRequest{
			Creator:      creator,
			Aim:          aim,
			Details:      details,
			Reason:       reason,
			Type:         typ,
			Tags:         tags,
			Journey:      &journey,
			IsAchieved:   false,
			IsPrivate:    false,
			Likes:        []string{},
			SimilarGoals: tags,
			Inspiration:  "",
			Watchers:     []string{},
		},
	)
	if err != nil {
		log.Println("error in proto")
		return err.Error(), err
	}

	log.Println(r.Goal)
	return r.Goal, nil
}

// Put ...
func (c *Client) Put(ctx context.Context, id, creator, aim, reason, details, typ, inspiration string,
	tags, likes, similarGoals, watchers []string,
	isAchieved, isPrivate bool,
	journeyType, journeyDetails string,
	journeyIsStarted, journeyIsInProgress, journeyIsComplete bool,
	journeySteps []string) (string, error) {
	var journey = pb.Journey{}

	journey.Typ = journeyType
	journey.IsInProgress = journeyIsInProgress
	journey.IsStarted = journeyIsStarted
	journey.IsComplete = journeyIsComplete
	journey.Details = journeyDetails
	journey.Steps = journeySteps

	r, err := c.service.PutGoal(
		ctx,
		&pb.PutGoalRequest{
			Id:           id,
			Creator:      creator,
			Aim:          aim,
			Details:      details,
			Reason:       reason,
			Type:         typ,
			Tags:         tags,
			Journey:      &journey,
			IsAchieved:   isAchieved,
			IsPrivate:    isPrivate,
			Likes:        likes,
			SimilarGoals: similarGoals,
			Inspiration:  inspiration,
			Watchers:     watchers,
		},
	)
	if err != nil {
		return "", err
	}
	return r.Goal, nil
}

// Get ...
func (c *Client) Get(ctx context.Context, id string, userID uint64) (*domain.Goal, error) {
	var goal = domain.Goal{}
	var journey = pb.Journey{}
	var participants []bson.ObjectId
	var watchers []bson.ObjectId
	var likes []bson.ObjectId

	r, err := c.service.GetGoal(
		ctx,
		&pb.GetGoalRequest{
			Id:     id,
			UserID: userID,
		},
	)
	if err != nil {
		return nil, err
	}

	for _, f := range r.Goal.Participants {
		participants = append(participants, bson.ObjectIdHex(f))
	}

	for _, f := range r.Goal.Watchers {
		watchers = append(watchers, bson.ObjectIdHex(f))
	}

	for _, f := range r.Goal.Likes {
		likes = append(likes, bson.ObjectIdHex(f))
	}

	journey = *r.Goal.GoalJourney
	goal.Journey.Details = journey.Details
	goal.Journey.IsComplete = journey.IsComplete
	goal.Journey.IsInProgress = journey.IsInProgress
	goal.Journey.IsStarted = journey.IsStarted
	goal.Journey.Progress = journey.Progress
	goal.Journey.Steps = journey.Steps

	return &domain.Goal{
		ID:           bson.ObjectIdHex(r.Goal.Id),
		Creator:      bson.ObjectIdHex(r.Goal.Creator),
		Aim:          r.Goal.Aim,
		Details:      r.Goal.Details,
		Reason:       r.Goal.Reason,
		Type:         r.Goal.Type,
		Likes:        likes,
		Participants: participants,
		Watchers:     watchers,
		IsAchieved:   r.Goal.IsAchieved,
		IsPrivate:    r.Goal.IsPrivate,
		SimilarGoals: r.Goal.SimilarGoals,
		Journey:      goal.Journey,
		Tags:         r.Goal.Tags,
	}, nil
}

// GetMultiple is used to get the list of specified goals
func (c *Client) GetMultiple(ctx context.Context, skip, take uint64) ([]domain.Goal, error) {
	var goal = domain.Goal{}
	var journey = pb.Journey{}
	var participants []bson.ObjectId
	var watchers []bson.ObjectId
	var likes []bson.ObjectId

	r, err := c.service.GetMultipleGoals(
		ctx,
		&pb.GetMultipleGoalsRequest{
			Skip: skip,
			Take: take,
		},
	)
	if err != nil {
		return nil, err
	}

	goals := []domain.Goal{}
	for _, a := range r.Goals {
		for _, f := range a.Participants {
			participants = append(participants, bson.ObjectIdHex(f))
		}

		for _, f := range a.Watchers {
			watchers = append(watchers, bson.ObjectIdHex(f))
		}

		for _, f := range a.Likes {
			likes = append(likes, bson.ObjectIdHex(f))
		}

		journey = *a.GoalJourney
		goal.Journey.Details = journey.Details
		goal.Journey.IsComplete = journey.IsComplete
		goal.Journey.IsInProgress = journey.IsInProgress
		goal.Journey.IsStarted = journey.IsStarted
		goal.Journey.Progress = journey.Progress
		goal.Journey.Steps = journey.Steps

		goals = append(goals, domain.Goal{
			ID:           bson.ObjectIdHex(a.Id),
			Creator:      bson.ObjectIdHex(a.Creator),
			Aim:          a.Aim,
			Details:      a.Details,
			Reason:       a.Reason,
			Type:         a.Type,
			Likes:        likes,
			Participants: participants,
			Watchers:     watchers,
			IsAchieved:   a.IsAchieved,
			IsPrivate:    a.IsPrivate,
			SimilarGoals: a.SimilarGoals,
			Journey:      goal.Journey,
			Tags:         a.Tags,
		})
	}
	return goals, nil
}

// Delete removes a goal with passed identifier
func (c *Client) Delete(ctx context.Context, id string) (string, error) {
	r, err := c.service.DeleteGoal(
		ctx,
		&pb.DeleteGoalRequest{Id: id},
	)
	if err != nil {
		return "", err
	}
	return r.Id, nil
}
