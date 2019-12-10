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
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/hackerrithm/blackfox/services/backend/auth/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/auth/pkg/engine"
	pb "github.com/hackerrithm/blackfox/services/backend/auth/pkg/model"
	profile "github.com/hackerrithm/blackfox/services/backend/profile/cmd/profile/client"
)

type grpcServer struct {
	service       engine.Auth
	profileClient *profile.Client
}

// ListenGRPC ...
func ListenGRPC(s engine.Auth, profileURL string, port int) error {
	profileClient, err := profile.NewClient(profileURL)
	if err != nil {
		profileClient.Close()
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterAuthServiceServer(serv, &grpcServer{s, profileClient})
	reflection.Register(serv)
	return serv.Serve(lis)
}

// Check is for health checking.
func (s *grpcServer) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (s *grpcServer) UserRegister(ctx context.Context, r *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	var user domain.Register
	user.UserName = r.Register.Username
	user.FirstName = r.Register.Firstname
	user.LastName = r.Register.Lastname
	user.Password = r.Register.Password
	user.Gender = r.Register.Gender
	user.Status = r.Register.Status

	a, err := s.service.RegisterUser(ctx, user)
	if err != nil {
		return nil, err
	}

	profileObject, err := s.profileClient.Post(ctx, user.UserName, "1", "not much here", make([]string, 0), make([]string, 0), 1)
	if err != nil {
		return nil, err
	}

	fmt.Println(profileObject, " profile created")

	return &pb.UserRegisterResponse{
		User: &pb.User{
			Username:     a.Username,
			Password:     a.Password,
			Firstname:    a.Firstname,
			Lastname:     a.Lastname,
			Gender:       a.Gender,
			Emailaddress: a.EmailAddress,
			Name:         a.Name,
			Status:       a.Status,
			Type:         a.Type,
			Middlename:   a.Gender,
		},
	}, nil
}

func (s *grpcServer) UserLogin(ctx context.Context, r *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	var auth domain.Login
	auth.UserName = r.Username
	auth.Password = r.Password

	a, err := s.service.LoginUser(ctx, auth)
	if err != nil {
		return nil, err
	}

	res, err := s.service.GenerateToken(ctx, a.ID.Hex())
	if err != nil {
		return &pb.UserLoginResponse{}, err
	}

	return &pb.UserLoginResponse{
		Token: res,
	}, nil
}

func (s *grpcServer) GenerateToken(ctx context.Context, r *pb.GenerateTokenRequest) (*pb.GenerateTokenResponse, error) {
	var ID = r.Token
	b, err := s.service.GenerateToken(ctx, ID)
	if err != nil {
		return &pb.GenerateTokenResponse{}, err
	}

	return &pb.GenerateTokenResponse{
		Result: b,
	}, nil
}

func (s *grpcServer) GetUserFromToken(ctx context.Context, r *pb.GetUserFromTokenRequest) (*pb.GetUserFromTokenResponse, error) {
	log.Println("got in search for GetUserFromToken grpc method-----------")

	var ID = r.Token
	log.Println("------------------- ID VAL: ", ID)
	claims, err := s.service.ParseToken(ctx, ID)
	if err != nil {
		return &pb.GetUserFromTokenResponse{}, err
	}

	// if err, ok := claims["type"].(float64); ok != true {
	// 	return &pb.GenerateTokenResponse{}, err
	// }

	userid, ok := claims["userid"].(string)
	if !ok {
		return &pb.GetUserFromTokenResponse{}, fmt.Errorf("user id can't get from token claims: %v", claims)
	}
	log.Println("------------------- ID VAL:- USERID ", userid)
	return &pb.GetUserFromTokenResponse{
		Result: userid,
	}, nil
}
