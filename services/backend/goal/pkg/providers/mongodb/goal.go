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

package mongodb

import (
	"log"

	"golang.org/x/net/context"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	cfg "github.com/hackerrithm/blackfox/services/backend/goal/configs"
	"github.com/hackerrithm/blackfox/services/backend/goal/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/goal/pkg/engine"
)

type (
	goalRepository struct {
		session *mgo.Session
	}
)

const (
	goalCollection = "goal"
)

var (
	config cfg.Config
)

func newGoalRepository(session *mgo.Session) engine.GoalRepository {
	return &goalRepository{session}
}

func (r goalRepository) Insert(c context.Context, p domain.Goal, j domain.Journey) error {
	log.Println("arrived at repo")
	s := r.session.Clone()
	defer s.Close()

	var goal domain.Goal
	var journey domain.Journey

	journey.Details = j.Details
	journey.DueDate = j.DueDate
	journey.IsComplete = j.IsComplete
	journey.IsInProgress = j.IsInProgress
	journey.IsStarted = j.IsStarted
	journey.Progress = j.Progress
	journey.StartDate = j.StartDate
	journey.Steps = j.Steps
	journey.Type = j.Type

	goal.Aim = p.Aim
	goal.Creator = p.Creator
	goal.Details = p.Details
	goal.Inspiration = p.Inspiration
	goal.IsAchieved = p.IsAchieved
	goal.IsPrivate = p.IsPrivate
	goal.Journey = p.Journey
	goal.Type = p.Type
	goal.Likes = p.Likes
	goal.Tags = p.Tags
	goal.SimilarGoals = p.SimilarGoals
	goal.Watchers = p.Watchers
	goal.Participants = p.Participants
	goal.Reason = p.Reason
	goal.Date = p.Date

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Insert(&goal)
	if err != nil {
		log.Println("insert error ", err)
	}

	return nil
}

func (r goalRepository) Update(c context.Context, p domain.Goal, j domain.Journey, id string) error {
	s := r.session.Clone()
	defer s.Close()

	var goal domain.Goal
	var journey domain.Journey

	journey.Details = j.Details
	journey.DueDate = j.DueDate
	journey.IsComplete = j.IsComplete
	journey.IsInProgress = j.IsInProgress
	journey.IsStarted = j.IsStarted
	journey.Progress = j.Progress
	journey.StartDate = j.StartDate
	journey.Steps = j.Steps
	journey.Type = j.Type

	goal.Aim = p.Aim
	goal.Creator = p.Creator
	goal.Details = p.Details
	goal.Inspiration = p.Inspiration
	goal.IsAchieved = p.IsAchieved
	goal.IsPrivate = p.IsPrivate
	goal.Journey = p.Journey
	goal.Type = p.Type
	goal.Likes = p.Likes
	goal.Tags = p.Tags
	goal.SimilarGoals = p.SimilarGoals
	goal.Watchers = p.Watchers
	goal.Participants = p.Participants
	goal.Reason = p.Reason
	goal.Date = p.Date

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Update(bson.M{"_id": toBSON(id)}, p)
	if err != nil {
		return err
	}

	return nil
}

func (r goalRepository) Query(c context.Context, query *engine.Query) []*domain.Goal {
	s := r.session.Clone()
	defer s.Close()

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	p := []*domain.Goal{}
	q := translateQuery(col, query)
	q.All(&p)

	return nil
}

func (r goalRepository) FindOne(c context.Context, id string) (*domain.Goal, error) {
	s := r.session.Clone()
	defer s.Close()

	var goal *domain.Goal
	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"_id": toBSON(id)}).One(&goal)

	if err != nil {
		return nil, nil
	}

	return goal, nil
}

// ListAllGoals used for finding all user goals
// by the passed skip and take parameters
func (r goalRepository) ListAllGoals(ctx context.Context, skip uint64, take uint64) ([]domain.Goal, error) {
	s := r.session.Clone()
	defer s.Close()
	log.Println("got in repo")

	var goals []domain.Goal

	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(nil).All(&goals)
	if err != nil {
		return nil, err
	}

	return goals, nil
}

func (r goalRepository) Remove(c context.Context, id string) (string, error) {
	s := r.session.Clone()
	defer s.Close()

	err := s.DB(config.MongoDB).C(config.MongoCollection).Remove(bson.M{"_id": toBSON(id)}) //.One(&goal)

	if err != nil {
		return "", nil
	}

	return "", nil
}
