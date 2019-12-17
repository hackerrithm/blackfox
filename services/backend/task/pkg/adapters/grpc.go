//go:generate protoc ./task.proto --go_out=plugins=grpc:./pb
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
	"github.com/hackerrithm/blackfox/services/backend/task/pkg/engine"
	pb "github.com/hackerrithm/blackfox/services/backend/task/pkg/model"
)

type grpcServer struct {
	service engine.Task
	profile *userProfile.Client
}

// ListenGRPC ...
func ListenGRPC(s engine.Task, userURL string, port int) error {
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
	pb.RegisterTaskServiceServer(serv, &grpcServer{s, profileClient})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostTask(ctx context.Context, r *pb.PostTaskRequest) (*pb.PostTaskResponse, error) {
	err := s.service.Insert(ctx, r.Text)
	if err != nil {
		log.Println("here error in task method")
		return nil, err
	}
	return &pb.PostTaskResponse{
		Task: "inserted",
	}, nil
}

func (s *grpcServer) PutTask(ctx context.Context, r *pb.PutTaskRequest) (*pb.PutTaskResponse, error) {
	err := s.service.Update(ctx, r.Id, r.Text)
	if err != nil {
		return nil, err
	}
	return &pb.PutTaskResponse{
		Task: "updated",
	}, nil
}

func (s *grpcServer) GetTask(ctx context.Context, r *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	a, err := s.service.FindOne(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetTaskResponse{
		Task: &pb.Task{
			Id:   a.ID,
			Text: a.Text,
		},
	}, nil
}

func (s *grpcServer) GetMultipleTask(ctx context.Context, r *pb.GetMultipleTaskRequest) (*pb.GetMultipleTaskResponse, error) {
	res, err := s.service.ListAllTasks(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	tasks := []*pb.Task{}
	for _, p := range *res {
		tasks = append(
			tasks,
			&pb.Task{
				Id:   p.ID,
				Text: p.Text,
			},
		)
	}

	return &pb.GetMultipleTaskResponse{
		Tasks: tasks,
	}, nil
}

func (s *grpcServer) DeleteTask(ctx context.Context, r *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	a, err := s.service.Remove(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteTaskResponse{
		Id: uint32(a),
	}, nil
}
