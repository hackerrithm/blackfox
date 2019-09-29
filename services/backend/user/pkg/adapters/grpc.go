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

	profile "github.com/hackerrithm/blackfox/services/backend/profile/cmd/profile/client"
	"github.com/hackerrithm/blackfox/services/backend/user/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/user/pkg/engine"
	pb "github.com/hackerrithm/blackfox/services/backend/user/pkg/model"
)

type grpcServer struct {
	service       engine.User
	profileClient *profile.Client
}

// ListenGRPC ...
func ListenGRPC(s engine.User, profileURL string, port int) error {
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
	pb.RegisterUserServiceServer(serv, &grpcServer{s, profileClient})
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
	var user domain.Login
	user.UserName = r.Username
	user.Password = r.Password

	a, err := s.service.LoginUser(ctx, user)
	if err != nil {
		return nil, err
	}

	res, err := s.service.GenerateToken(ctx, a.ID.Hex())
	if err != nil {
		return &pb.UserLoginResponse{}, err
	}

	return &pb.UserLoginResponse{
		Token: res["token"].(string),
	}, nil
}

func (s *grpcServer) GetAllUsers(ctx context.Context, r *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	var uname, pwd, lname, fname []string
	var q = pb.User{}

	a, err := s.service.ListAllUsers(ctx, 0, 0)
	if err != nil {
		return nil, err
	}
	for index := 0; index < len(a); index++ {
		q.Username = a[index].Username
		q.Password = a[index].Password
		q.Firstname = a[index].Firstname
		q.Lastname = a[index].Lastname
		q.Gender = a[index].Gender
		q.Name = a[index].Name
		q.Status = a[index].Status
		q.Emailaddress = a[index].EmailAddress

		uname = append(uname, q.Username)
		pwd = append(pwd, q.Password)
		lname = append(lname, q.Lastname)
		fname = append(fname, q.Firstname)

	}
	return &pb.GetAllUsersResponse{
		Username:  uname,
		Password:  pwd,
		Firstname: fname,
		Lastname:  lname,
	}, nil
}

func (s *grpcServer) GetUser(ctx context.Context, r *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Println("in grpc with ", r.Id)
	a, err := s.service.GetUserByID(ctx, r.Id)
	if err != nil {
		return &pb.GetUserResponse{}, err
	}

	var q = pb.User{}
	q.Username = a.Username
	q.Password = a.Password
	q.Id = a.ID.Hex()
	q.Gender = a.Gender
	q.Name = a.Name
	q.Status = a.Status
	q.Lastname = a.Lastname
	q.Firstname = a.Firstname

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:           q.Id,
			Username:     q.Username,
			Password:     q.Password,
			Firstname:    q.Firstname,
			Lastname:     q.Lastname,
			Emailaddress: q.Emailaddress,
			Gender:       q.Gender,
		},
	}, nil
}

func (s *grpcServer) GetUserByUserName(ctx context.Context, r *pb.GetUserByUserNameRequest) (*pb.GetUserByUserNameResponse, error) {
	a, err := s.service.GetUserByUserName(ctx, r.Username)
	if err != nil {
		return &pb.GetUserByUserNameResponse{}, err
	}
	var q = pb.User{}
	q.Username = a.Username
	q.Password = a.Password
	q.Id = a.ID.Hex()
	q.Gender = a.Gender
	q.Name = a.Name
	q.Status = a.Status
	q.Lastname = a.Lastname
	q.Firstname = a.Firstname
	q.Emailaddress = a.EmailAddress

	return &pb.GetUserByUserNameResponse{
		User: &pb.User{
			Id:           q.Id,
			Username:     q.Username,
			Password:     q.Password,
			Firstname:    q.Firstname,
			Lastname:     q.Lastname,
			Emailaddress: q.Emailaddress,
			Gender:       q.Gender,
		},
	}, nil
}

func (s *grpcServer) GetUserByEmailAddress(ctx context.Context, r *pb.GetUserByEmailAddressRequest) (*pb.GetUserByEmailAddressResponse, error) {
	a, err := s.service.GetUserByEmailAddress(ctx, r.Email)
	if err != nil {
		return &pb.GetUserByEmailAddressResponse{}, err
	}
	var q = pb.User{}
	q.Username = a.Username
	q.Password = a.Password
	q.Id = a.ID.Hex()
	q.Gender = a.Gender
	q.Name = a.Name
	q.Status = a.Status
	q.Lastname = a.Lastname
	q.Firstname = a.Firstname
	q.Emailaddress = a.EmailAddress

	return &pb.GetUserByEmailAddressResponse{
		User: &pb.User{
			Id:           q.Id,
			Username:     q.Username,
			Password:     q.Password,
			Firstname:    q.Firstname,
			Lastname:     q.Lastname,
			Emailaddress: q.Emailaddress,
			Gender:       q.Gender,
		},
	}, nil
}

func (s *grpcServer) DeleteUser(ctx context.Context, r *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	var ID = r.Id
	b, err := s.service.DeleteUserAccount(ctx, ID)
	if err != nil {
		return &pb.DeleteUserResponse{}, err
	}

	return &pb.DeleteUserResponse{
		IsDeleted: b,
	}, nil
}
