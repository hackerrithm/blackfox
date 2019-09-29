//go:generate protoc ./geography.proto --go_out=plugins=grpc:./pb
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

	"github.com/hackerrithm/blackfox/services/backend/geography/pkg/engine"
	pb "github.com/hackerrithm/blackfox/services/backend/geography/pkg/model"
)

type grpcServer struct {
	service engine.Geography
}

// ListenGRPC ...
func ListenGRPC(s engine.Geography, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterGeographyServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) GetLocationDistance(ctx context.Context, r *pb.LocationDistanceRequest) (*pb.LocationDistanceResponse, error) {
	res, err := s.service.GetLocationDistance(ctx, r.Longitude, r.Latitude)
	if err != nil {
		log.Println("here error in geography method")
		return nil, err
	}
	log.Println("res:", res)
	return &pb.LocationDistanceResponse{
		Geography: res,
	}, nil
}
