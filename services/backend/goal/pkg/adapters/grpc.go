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

package adapters

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/hackerrithm/blackfox/services/backend/goal/pkg/engine"
	pb "github.com/hackerrithm/blackfox/services/backend/goal/pkg/model"
)

type grpcServer struct {
	service engine.Goal
	// profile *userProfile.Client
}

// ListenGRPC ...
func ListenGRPC(s engine.Goal, port int) error {
	// profileClient, err := userProfile.NewClient(userURL)
	// if err != nil {
	// 	profileClient.Close()
	// 	return err
	// }

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		// profileClient.Close()
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterGoalServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostGoal(ctx context.Context, r *pb.PostGoalRequest) (*pb.PostGoalResponse, error) {
	log.Println("got in server")
	err := s.service.Insert(ctx, r.Creator, r.Aim, r.Reason, r.Details, r.Type, r.Tags)
	if err != nil {
		log.Println("here error in goal method")
		return nil, err
	}
	return &pb.PostGoalResponse{
		Goal: "inserted",
	}, nil
}

func (s *grpcServer) PutGoal(ctx context.Context, r *pb.PutGoalRequest) (*pb.PutGoalResponse, error) {
	err := s.service.Update(ctx, r.Id, r.Creator, r.Aim, r.Reason, r.Details, r.Type, r.Tags)
	if err != nil {
		return nil, err
	}
	return &pb.PutGoalResponse{
		Goal: "updated",
	}, nil
}

func (s *grpcServer) GetGoal(ctx context.Context, r *pb.GetGoalRequest) (*pb.GetGoalResponse, error) {
	a, err := s.service.FindOne(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	var participants []string
	var watchers []string
	var likes []string

	for _, f := range a.Participants {
		participants = append(participants, f.Hex())
	}

	for _, f := range a.Watchers {
		watchers = append(watchers, f.Hex())
	}

	for _, f := range a.Likes {
		likes = append(likes, f.Hex())
	}

	var journey = pb.Journey{}
	journey.Details = a.Journey.Details
	// journey.DueDate = a.Journey.DueDate // to be dealt with due to inability to convert
	journey.IsComplete = a.Journey.IsComplete
	journey.IsInProgress = a.Journey.IsInProgress
	journey.IsStarted = a.Journey.IsStarted
	journey.Progress = a.Journey.Progress
	// journey.StartDate = a.Journey.StartDate // to be dealt with due to inability to convert
	journey.Steps = a.Journey.Steps
	// journey.Type = a.Journey.Type // to be dealt with due to inability to convert

	return &pb.GetGoalResponse{
		Goal: &pb.Goal{
			Id:           a.ID.Hex(),
			Creator:      a.Creator.Hex(),
			Aim:          a.Aim,
			Inspiration:  a.Inspiration,
			IsAchieved:   a.IsAchieved,
			IsPrivate:    a.IsPrivate,
			Details:      a.Details,
			SimilarGoals: a.SimilarGoals,
			Type:         a.Type,
			Likes:        likes,
			Watchers:     watchers,
			GoalJourney:  &journey,
			Participants: participants,
			Tags:         a.Tags,
		},
	}, nil
}

func (s *grpcServer) GetMultipleGoals(ctx context.Context, r *pb.GetMultipleGoalsRequest) (*pb.GetMultipleGoalsResponse, error) {
	log.Println("got to grpc ")
	var participants []string
	var watchers []string
	var likes []string

	res, err := s.service.ListAllGoals(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	goals := []*pb.Goal{}
	for _, p := range *res {

		for _, f := range p.Participants {
			participants = append(participants, f.Hex())
		}

		for _, f := range p.Watchers {
			watchers = append(watchers, f.Hex())
		}

		for _, f := range p.Likes {
			likes = append(likes, f.Hex())
		}

		var journey = pb.Journey{}
		journey.Details = p.Journey.Details
		// journey.DueDate = p.Journey.DueDate // to be dealt with due to inability to convert
		journey.IsComplete = p.Journey.IsComplete
		journey.IsInProgress = p.Journey.IsInProgress
		journey.IsStarted = p.Journey.IsStarted
		journey.Progress = p.Journey.Progress
		// journey.StartDate = p.Journey.StartDate // to be dealt with due to inability to convert
		journey.Steps = p.Journey.Steps
		// journey.Type = p.Journey.Type // to be dealt with due to inability to convert

		goals = append(
			goals,
			&pb.Goal{
				Id:           p.ID.Hex(),
				Creator:      p.Creator.Hex(),
				Aim:          p.Aim,
				Inspiration:  p.Inspiration,
				IsAchieved:   p.IsAchieved,
				IsPrivate:    p.IsPrivate,
				Details:      p.Details,
				SimilarGoals: p.SimilarGoals,
				Type:         p.Type,
				Likes:        likes,
				Watchers:     watchers,
				GoalJourney:  &journey,
				Participants: participants,
				Tags:         p.Tags,
			},
		)
	}

	return &pb.GetMultipleGoalsResponse{
		Goals: goals,
	}, nil
}

func (s *grpcServer) DeleteGoal(ctx context.Context, r *pb.DeleteGoalRequest) (*pb.DeleteGoalResponse, error) {
	a, err := s.service.Remove(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteGoalResponse{
		Id: a,
	}, nil
}
