//go:generate protoc ./space.proto --go_out=plugins=grpc:./pb
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
	"github.com/hackerrithm/blackfox/services/backend/space/pkg/engine"
	pb "github.com/hackerrithm/blackfox/services/backend/space/pkg/model"
)

type grpcServer struct {
	service engine.Space
	profile *userProfile.Client
}

// ListenGRPC ...
func ListenGRPC(s engine.Space, userURL string, port int) error {
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
	pb.RegisterSpaceServiceServer(serv, &grpcServer{s, profileClient})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostSpace(ctx context.Context, r *pb.PostSpaceRequest) (*pb.PostSpaceResponse, error) {
	err := s.service.Insert(ctx, r.Creator, r.Topic, r.Details, r.Description, r.Type, r.Managers, r.Followers, r.Tags)
	if err != nil {
		log.Println("here error in space method")
		return nil, err
	}
	return &pb.PostSpaceResponse{
		Space: "inserted",
	}, nil
}

func (s *grpcServer) PutSpace(ctx context.Context, r *pb.PutSpaceRequest) (*pb.PutSpaceResponse, error) {
	err := s.service.Update(ctx, r.Id, r.Creator, r.Topic, r.Details, r.Description, r.Type, r.Managers, r.Followers, r.Tags)
	if err != nil {
		return nil, err
	}
	return &pb.PutSpaceResponse{
		Space: "updated",
	}, nil
}

func (s *grpcServer) GetSpace(ctx context.Context, r *pb.GetSpaceRequest) (*pb.GetSpaceResponse, error) {
	a, err := s.service.FindOne(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	var mngs, fllwrs []string
	for _, m := range a.Managers {
		mngs = append(mngs, m.Hex())
	}
	for _, f := range a.Followers {
		fllwrs = append(fllwrs, f.Hex())
	}

	return &pb.GetSpaceResponse{
		Space: &pb.Space{
			Id:          a.ID.Hex(),
			Creator:     a.Creator.Hex(),
			Topic:       a.Topic,
			Details:     a.Details,
			Description: a.Description,
			Type:        a.Type,
			Followers:   fllwrs,
			Managers:    mngs,
			Tags:        a.Tags,
		},
	}, nil
}

func (s *grpcServer) GetMultipleSpaces(ctx context.Context, r *pb.GetMultipleSpacesRequest) (*pb.GetMultipleSpacesResponse, error) {
	log.Println("got to grpc ")
	var mngs, fllwrs []string

	res, err := s.service.ListAllSpaces(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	spaces := []*pb.Space{}
	for _, p := range *res {
		for _, m := range p.Managers {
			mngs = append(mngs, m.Hex())
		}
		for _, f := range p.Followers {
			fllwrs = append(fllwrs, f.Hex())
		}
		spaces = append(
			spaces,
			&pb.Space{
				Id:          p.ID.Hex(),
				Creator:     p.Creator.Hex(),
				Topic:       p.Topic,
				Details:     p.Details,
				Description: p.Description,
				Type:        p.Type,
				Followers:   fllwrs,
				Managers:    mngs,
				Tags:        p.Tags,
			},
		)
	}

	return &pb.GetMultipleSpacesResponse{
		Spaces: spaces,
	}, nil
}

func (s *grpcServer) DeleteSpace(ctx context.Context, r *pb.DeleteSpaceRequest) (*pb.DeleteSpaceResponse, error) {
	a, err := s.service.Remove(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteSpaceResponse{
		Id: a,
	}, nil
}
