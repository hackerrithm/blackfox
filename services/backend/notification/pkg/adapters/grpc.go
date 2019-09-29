//go:generate protoc ./notification.proto --go_out=plugins=grpc:./pb
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

	"github.com/hackerrithm/blackfox/services/backend/notification/pkg/engine"
	pb "github.com/hackerrithm/blackfox/services/backend/notification/pkg/model"
)

type grpcServer struct {
	service engine.Notification
}

// ListenGRPC ...
func ListenGRPC(s engine.Notification, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterNotificationServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) SendNotification(ctx context.Context, r *pb.SendNotificationRequest) (*pb.SendNotificationResponse, error) {
	res, err := s.service.SendNotification(ctx, r.Service, r.Topic, r.Body)
	if err != nil {
		log.Println("here error in notifications method")
		return nil, err
	}
	log.Println("res:", res)
	return &pb.SendNotificationResponse{
		Notification: &pb.Notification{
			Service: res.Service,
			Topic:   res.Body,
			Body:    res.Body,
		},
	}, nil
}
