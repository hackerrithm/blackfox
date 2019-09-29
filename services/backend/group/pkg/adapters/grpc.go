//go:generate protoc ./group.proto --go_out=plugins=grpc:./pb
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

	"github.com/hackerrithm/blackfox/services/backend/group/pkg/engine"
	pb "github.com/hackerrithm/blackfox/services/backend/group/pkg/model"
	userProfile "github.com/hackerrithm/blackfox/services/backend/profile/cmd/profile/client"
)

type grpcServer struct {
	service engine.Group
	profile *userProfile.Client
}

// ListenGRPC ...
func ListenGRPC(s engine.Group, userURL string, port int) error {
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
	pb.RegisterGroupServiceServer(serv, &grpcServer{s, profileClient})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostGroup(ctx context.Context, r *pb.PostGroupRequest) (*pb.PostGroupResponse, error) {
	err := s.service.Insert(ctx, r.Title, r.Details, r.Description, r.Type, r.People)
	if err != nil {
		log.Println("here error in group method")
		return nil, err
	}
	return &pb.PostGroupResponse{
		Group: "inserted",
	}, nil
}

func (s *grpcServer) PutGroup(ctx context.Context, r *pb.PutGroupRequest) (*pb.PutGroupResponse, error) {
	err := s.service.Update(ctx, r.Id, r.Title, r.Details, r.Description, r.Type, r.People)
	if err != nil {
		return nil, err
	}
	return &pb.PutGroupResponse{
		Group: "updated",
	}, nil
}

func (s *grpcServer) GetGroup(ctx context.Context, r *pb.GetGroupRequest) (*pb.GetGroupResponse, error) {
	a, err := s.service.FindOne(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	var people []string
	for _, m := range a.People {
		people = append(people, m.Hex())
	}

	return &pb.GetGroupResponse{
		Group: &pb.Group{
			Id:          a.ID.Hex(),
			Title:       a.Title,
			Details:     a.Details,
			Description: a.Description,
			Type:        a.Type,
			People:      people,
		},
	}, nil
}

func (s *grpcServer) GetMultipleGroups(ctx context.Context, r *pb.GetMultipleGroupsRequest) (*pb.GetMultipleGroupsResponse, error) {
	log.Println("got to grpc ")
	var people []string

	res, err := s.service.ListAllGroups(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	groups := []*pb.Group{}
	for _, p := range *res {
		for _, m := range p.People {
			people = append(people, m.Hex())
		}
		groups = append(
			groups,
			&pb.Group{
				Id:          p.ID.Hex(),
				Title:       p.Title,
				Details:     p.Details,
				Description: p.Description,
				Type:        p.Type,
				People:      people,
			},
		)
	}

	return &pb.GetMultipleGroupsResponse{
		Groups: groups,
	}, nil
}

func (s *grpcServer) DeleteGroup(ctx context.Context, r *pb.DeleteGroupRequest) (*pb.DeleteGroupResponse, error) {
	a, err := s.service.Remove(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteGroupResponse{
		Id: a,
	}, nil
}
