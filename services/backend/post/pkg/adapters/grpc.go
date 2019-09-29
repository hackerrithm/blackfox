//go:generate protoc ./post.proto --go_out=plugins=grpc:./pb
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

	"github.com/hackerrithm/blackfox/services/backend/post/pkg/engine"
	pb "github.com/hackerrithm/blackfox/services/backend/post/pkg/model"
	userProfile "github.com/hackerrithm/blackfox/services/backend/profile/cmd/profile/client"
)

type grpcServer struct {
	service engine.Post
	profile *userProfile.Client
}

// ListenGRPC ...
func ListenGRPC(s engine.Post, userURL string, port int) error {
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
	pb.RegisterPostServiceServer(serv, &grpcServer{s, profileClient})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostPost(ctx context.Context, r *pb.PostPostRequest) (*pb.PostPostResponse, error) {
	// fileName, err := FileUpload(ctx.Writer, ctx.Request)
	// if err != nil {
	// 	log.Println("error bya")
	// }

	img := pb.Image{}
	img.Name = r.ContentPhoto.Name

	err := s.service.Insert(ctx, r.Author, r.Topic, r.Category, r.ContentText, img.Name)
	if err != nil {
		log.Println("here error in post method")
		return nil, err
	}
	return &pb.PostPostResponse{
		Post: "inserted",
	}, nil
}

func (s *grpcServer) PutPost(ctx context.Context, r *pb.PutPostRequest) (*pb.PutPostResponse, error) {
	img := pb.Image{}
	img.Name = r.ContentPhoto.Name

	err := s.service.Update(ctx, r.Author, r.Topic, r.Category, r.ContentText, img.Name, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.PutPostResponse{
		Post: "updated",
	}, nil
}

func (s *grpcServer) GetPost(ctx context.Context, r *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	a, err := s.service.FindOne(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	var img = pb.Image{}
	img.Name = a.ContentPhoto.Name

	return &pb.GetPostResponse{
		Post: &pb.Post{
			Id:           a.ID.Hex(),
			Author:       a.Author.Hex(),
			Topic:        a.Topic,
			Category:     a.Category,
			ContentText:  a.ContentText,
			ContentPhoto: &img,
		},
	}, nil
}

func (s *grpcServer) GetMultiplePosts(ctx context.Context, r *pb.GetMultiplePostsRequest) (*pb.GetMultiplePostsResponse, error) {
	res, err := s.service.ListAllPosts(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	posts := []*pb.Post{}
	for _, p := range *res {
		var img = pb.Image{}
		img.Name = p.ContentPhoto.Name
		posts = append(
			posts,
			&pb.Post{
				Id:           p.ID.Hex(),
				Author:       p.Author.Hex(),
				Topic:        p.Topic,
				Category:     p.Category,
				ContentText:  p.ContentText,
				ContentPhoto: &img,
			},
		)
	}

	return &pb.GetMultiplePostsResponse{
		Posts: posts,
	}, nil
}

func (s *grpcServer) DeletePost(ctx context.Context, r *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	a, err := s.service.Remove(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeletePostResponse{
		Id: a,
	}, nil
}
