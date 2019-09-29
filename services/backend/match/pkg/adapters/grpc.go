//go:generate protoc ./match.proto --go_out=plugins=grpc:./pb
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

	goal "github.com/hackerrithm/blackfox/services/backend/goal/cmd/goal/client"
	"github.com/hackerrithm/blackfox/services/backend/match/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/match/pkg/engine"
	pb "github.com/hackerrithm/blackfox/services/backend/match/pkg/model"
	userProfile "github.com/hackerrithm/blackfox/services/backend/profile/cmd/profile/client"
	user "github.com/hackerrithm/blackfox/services/backend/user/cmd/user/client"
)

type grpcServer struct {
	service engine.Match
	profile *userProfile.Client
	goal    *goal.Client
	user    *user.Client
}

// ListenGRPC ...
func ListenGRPC(s engine.Match, userProfileURL, goalURL, userURL string, port int) error {
	profileClient, err := userProfile.NewClient(userProfileURL)
	if err != nil {
		profileClient.Close()
		return err
	}

	goalClient, err := goal.NewClient(goalURL)
	if err != nil {
		profileClient.Close()
		goalClient.Close()
		return err
	}

	userClient, err := user.NewClient(userURL)
	if err != nil {
		profileClient.Close()
		goalClient.Close()
		userClient.Close()
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		profileClient.Close()
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterMatchServiceServer(serv, &grpcServer{s, profileClient, goalClient, userClient})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostMatch(ctx context.Context, r *pb.PostMatchRequest) (*pb.PostMatchResponse, error) {
	err := s.service.Insert(ctx, r.PersonID, r.Details, r.Description, r.Type, r.Similarities)
	if err != nil {
		log.Println("here error in match method")
		return nil, err
	}
	return &pb.PostMatchResponse{
		Match: "inserted",
	}, nil
}

func (s *grpcServer) PutMatch(ctx context.Context, r *pb.PutMatchRequest) (*pb.PutMatchResponse, error) {
	err := s.service.Update(ctx, r.Id, r.PersonID, r.Details, r.Description, r.Type, r.Similarities)
	if err != nil {
		return nil, err
	}
	return &pb.PutMatchResponse{
		Match: "updated",
	}, nil
}

func (s *grpcServer) GetMatch(ctx context.Context, r *pb.GetMatchRequest) (*pb.GetMatchResponse, error) {
	a, err := s.service.FindOne(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	var smlrts []string
	for _, m := range a.Similarities {
		smlrts = append(smlrts, m)
	}

	return &pb.GetMatchResponse{
		Match: &pb.Match{
			Id:           a.ID.Hex(),
			PersonID:     a.PersonID.Hex(),
			Details:      a.Details,
			Description:  a.Description,
			Type:         a.Type,
			Similarities: smlrts,
		},
	}, nil
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func (s *grpcServer) GetMultipleMatches(ctx context.Context, r *pb.GetMultipleMatchesRequest) (*pb.GetMultipleMatchesResponse, error) {
	log.Println("got to grpc ")

	allRelevantGoals, err := s.goal.GetMultiple(ctx, 100, 100)
	if err != nil {
		return nil, err
	}

	requestSimilarities := []string{"love", "hate", "anger"}
	var matchingUsers = []string{""}

	for _, relevantGoal := range allRelevantGoals {
		if contains(requestSimilarities, relevantGoal.Aim) {
			matchingUsers = append(matchingUsers, relevantGoal.Creator.Hex())
		}
	}

	var matchedUsers []domain.MatchedUser
	var matchedUser domain.MatchedUser

	if len(matchingUsers) > 0 {
		for i := 0; i < len(matchingUsers); i++ {
			singleUser, err := s.user.GetUser(ctx, matchingUsers[i])
			if err != nil {
				return nil, err
			}
			matchedUser.EmailAddress = singleUser.EmailAddress
			matchedUser.Username = singleUser.Username
			matchedUser.Firstname = singleUser.Firstname
			matchedUser.Lastname = singleUser.Lastname
			matchedUser.Gender = singleUser.Gender
			matchedUser.ID = singleUser.ID
			matchedUser.Name = singleUser.Name

			matchedUsers = append(matchedUsers, matchedUser)
		}
	}

	matches := []*pb.MatchedUser{}
	for _, p := range matchedUsers {
		matches = append(
			matches,
			&pb.MatchedUser{
				Id:           p.ID.Hex(),
				Name:         p.Name,
				Username:     p.Username,
				Firstname:    p.Firstname,
				Lastname:     p.Lastname,
				Gender:       p.Gender,
				Type:         p.Type,
				EmailAddress: p.Gender,
			},
		)
	}

	return &pb.GetMultipleMatchesResponse{
		Matches: matches,
	}, nil
}

func (s *grpcServer) DeleteMatch(ctx context.Context, r *pb.DeleteMatchRequest) (*pb.DeleteMatchResponse, error) {
	a, err := s.service.Remove(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteMatchResponse{
		Id: a,
	}, nil
}
