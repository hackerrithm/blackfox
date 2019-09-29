//go:generate protoc ./profile.proto --go_out=plugins=grpc:./pb
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

	"github.com/hackerrithm/blackfox/services/backend/profile/pkg/engine"
	pb "github.com/hackerrithm/blackfox/services/backend/profile/pkg/model"
)

type grpcServer struct {
	service engine.Profile
}

// ListenGRPC ...
func ListenGRPC(s engine.Profile, userURL string, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterProfileServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) Post(ctx context.Context, r *pb.PostRequest) (*pb.PostResponse, error) {
	err := s.service.Insert(ctx, r.Username, r.Level, r.About, r.Followers, r.Following, r.Rings)
	if err != nil {
		log.Println("here error in profile method")
		return nil, err
	}
	return &pb.PostResponse{
		Profile: "inserted",
	}, nil
}

func (s *grpcServer) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
	err := s.service.Update(ctx, r.Id, r.Username, r.Level, r.About, r.Followers, r.Following, r.Rings)
	if err != nil {
		return nil, err
	}
	return &pb.PutResponse{
		Profile: "updated",
	}, nil
}

func (s *grpcServer) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	a, err := s.service.FindOne(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	var fllwrs []string
	var fllwing []string
	for _, f := range a.Followers {
		fllwrs = append(fllwrs, f.Hex())
	}

	for _, f := range a.Following {
		fllwing = append(fllwing, f.Hex())
	}

	var bkgimg = pb.Image{}
	var profimg = pb.Image{}
	bkgimg.Name = a.BackgroundImage.Name
	profimg.Name = a.ProfileImage.Name

	return &pb.GetResponse{
		Profile: &pb.Profile{
			Id:            a.ID.Hex(),
			Username:      a.UserName,
			About:         a.About,
			BackgroundImg: &bkgimg,
			ProfileImg:    &profimg,
			Followers:     fllwrs,
			Following:     fllwing,
			Rings:         a.Rings,
			Level:         a.Level,
		},
	}, nil
}

func (s *grpcServer) GetMultiple(ctx context.Context, r *pb.GetMultipleRequest) (*pb.GetMultipleResponse, error) {
	log.Println("got to grpc ")
	var fllwrs, fllwing []string

	res, err := s.service.ListAllProfiles(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	profiles := []*pb.Profile{}
	for _, p := range *res {
		var bkgimg = pb.Image{}
		var profimg = pb.Image{}
		bkgimg.Name = p.BackgroundImage.Name
		profimg.Name = p.ProfileImage.Name

		for _, f := range p.Followers {
			fllwrs = append(fllwrs, f.Hex())
		}

		for _, f := range p.Following {
			fllwing = append(fllwing, f.Hex())
		}

		profiles = append(
			profiles,
			&pb.Profile{
				Id:            p.ID.Hex(),
				Username:      p.UserName,
				About:         p.About,
				Rings:         p.Rings,
				Level:         p.Level,
				BackgroundImg: &bkgimg,
				ProfileImg:    &profimg,
				Followers:     fllwrs,
				Following:     fllwing,
			},
		)
	}

	return &pb.GetMultipleResponse{
		Profiles: profiles,
	}, nil
}

func (s *grpcServer) Delete(ctx context.Context, r *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	a, err := s.service.Remove(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{
		Id: a,
	}, nil
}
