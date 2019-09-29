//go:generate protoc ./reaction.proto --go_out=plugins=grpc:./pb
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

	userProfile "github.com/hackerrithm/blackfox/services/backend/profile/cmd/profile/client"
	"github.com/hackerrithm/blackfox/services/backend/reaction/pkg/engine"
	pb "github.com/hackerrithm/blackfox/services/backend/reaction/pkg/model"
)

type grpcServer struct {
	service engine.Reaction
	profile *userProfile.Client
}

// ListenGRPC ...
func ListenGRPC(s engine.Reaction, userURL string, port int) error {
	profileClient, err := userProfile.NewClient(userURL)
	if err != nil {
		profileClient.Close()
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		profileClient.Close()
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterReactionServiceServer(serv, &grpcServer{s, profileClient})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostReaction(ctx context.Context, r *pb.PostReactionRequest) (*pb.PostReactionResponse, error) {
	err := s.service.Insert(ctx, r.PersonID, r.Details, r.Description, r.Type, r.Similarities)
	if err != nil {
		log.Println("here error in reaction method")
		return nil, err
	}
	return &pb.PostReactionResponse{
		Reaction: "inserted",
	}, nil
}

func (s *grpcServer) PutReaction(ctx context.Context, r *pb.PutReactionRequest) (*pb.PutReactionResponse, error) {
	err := s.service.Update(ctx, r.Id, r.PersonID, r.Details, r.Description, r.Type, r.Similarities)
	if err != nil {
		return nil, err
	}
	return &pb.PutReactionResponse{
		Reaction: "updated",
	}, nil
}

func (s *grpcServer) GetReaction(ctx context.Context, r *pb.GetReactionRequest) (*pb.GetReactionResponse, error) {
	a, err := s.service.FindOne(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetReactionResponse{
		Reaction: &pb.Reaction{
			Id:          a.ID.Hex(),
			PersonID:    a.PersonID.Hex(),
			Details:     a.Details,
			Description: a.Description,
			Type:        a.Type,
			Tags:        a.Tags,
		},
	}, nil
}

func (s *grpcServer) GetMultipleReactions(ctx context.Context, r *pb.GetMultipleReactionsRequest) (*pb.GetMultipleReactionsResponse, error) {
	log.Println("got to grpc ")

	res, err := s.service.ListAllReactiones(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	reactions := []*pb.Reaction{}
	for _, p := range *res {
		reactions = append(
			reactions,
			&pb.Reaction{
				Id:          p.ID.Hex(),
				PersonID:    p.PersonID.Hex(),
				Details:     p.Details,
				Description: p.Description,
				Type:        p.Type,
			},
		)
	}

	return &pb.GetMultipleReactionsResponse{
		Reactions: reactions,
	}, nil
}

func (s *grpcServer) DeleteReaction(ctx context.Context, r *pb.DeleteReactionRequest) (*pb.DeleteReactionResponse, error) {
	a, err := s.service.Remove(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteReactionResponse{
		Id: a,
	}, nil
}
