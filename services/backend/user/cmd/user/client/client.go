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

package client

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"

	"github.com/hackerrithm/blackfox/services/backend/user/pkg/domain"
	pb "github.com/hackerrithm/blackfox/services/backend/user/pkg/model"
)

// Client act as our model
type Client struct {
	conn    *grpc.ClientConn
	service pb.UserServiceClient
}

// NewClient creates a new client instance
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("error in server !!!!!!!!!: %v", err)
	}

	c := pb.NewUserServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close closes connection pool
func (c *Client) Close() {
	c.conn.Close()
}

// LoginUser used to login a user
func (c *Client) LoginUser(ctx context.Context, username, password string) (*pb.UserLoginResponse, error) {
	r, err := c.service.UserLogin(
		ctx,
		&pb.UserLoginRequest{
			Username: username,
			Password: password,
		},
	)
	if err != nil {
		return nil, err
	}
	return &pb.UserLoginResponse{
		Token: r.Token,
	}, nil
}

// RegisterUser used to register a user
func (c *Client) RegisterUser(ctx context.Context, username, password, firstname, lastname, email, gender string) (string, error) {
	r, err := c.service.UserRegister(
		ctx,
		&pb.UserRegisterRequest{
			Register: &pb.Register{
				Username:  username,
				Password:  password,
				Firstname: firstname,
				Lastname:  lastname,
				Gender:    gender,
				Status:    "",
			},
		})
	if err != nil {
		return "", nil
	}
	log.Println(r.User)
	return "created", nil
}

// GetAllUsers ...
func (c *Client) GetAllUsers(ctx context.Context, ID string) ([]domain.User, error) {
	var userList = []domain.User{}
	var userCurrent = domain.User{}
	r, err := c.service.GetAllUsers(
		ctx,
		&pb.GetAllUsersRequest{
			Id: ID,
		})
	if err != nil {
		return []domain.User{}, err
	}

	for index := 0; index < len(r.Username); index++ {
		userCurrent.Username = r.Username[index]
		userCurrent.Password = r.Password[index]
		userList = append(userList, userCurrent)
	}

	return userList, nil
}

// GetUser ...
func (c *Client) GetUser(ctx context.Context, ID string) (*domain.User, error) {
	r, err := c.service.GetUser(
		ctx,
		&pb.GetUserRequest{
			Id: ID,
		})
	if err != nil {
		return &domain.User{}, err
	}
	return &domain.User{
		ID:           bson.ObjectId(r.User.Id),
		Username:     r.User.Username,
		Password:     r.User.Password,
		Firstname:    r.User.Firstname,
		Lastname:     r.User.Lastname,
		EmailAddress: r.User.Emailaddress,
		Gender:       r.User.Gender,
	}, nil
}

// GetUserByUserName ...
func (c *Client) GetUserByUserName(ctx context.Context, username string) (*domain.User, error) {
	r, err := c.service.GetUserByUserName(
		ctx,
		&pb.GetUserByUserNameRequest{
			Username: username,
		})
	if err != nil {
		return &domain.User{}, err
	}
	return &domain.User{
		ID:           bson.ObjectId(r.User.Id),
		Username:     r.User.Username,
		Password:     r.User.Password,
		Firstname:    r.User.Firstname,
		Lastname:     r.User.Lastname,
		EmailAddress: r.User.Emailaddress,
		Gender:       r.User.Gender,
	}, nil
}

// GetUserByEmailAddress ...
func (c *Client) GetUserByEmailAddress(ctx context.Context, email string) (*domain.User, error) {
	r, err := c.service.GetUserByEmailAddress(
		ctx,
		&pb.GetUserByEmailAddressRequest{
			Email: email,
		})
	if err != nil {
		return &domain.User{}, err
	}
	return &domain.User{
		ID:           bson.ObjectId(r.User.Id),
		Username:     r.User.Username,
		Password:     r.User.Password,
		Firstname:    r.User.Firstname,
		Lastname:     r.User.Lastname,
		EmailAddress: r.User.Emailaddress,
		Gender:       r.User.Gender,
	}, nil
}

// DeleteUser ...
func (c *Client) DeleteUser(ctx context.Context, ID string) (bool, error) {
	r, err := c.service.DeleteUser(
		ctx,
		&pb.DeleteUserRequest{
			Id: ID,
		})
	if err != nil {
		return r.IsDeleted, err
	}
	return r.IsDeleted, nil
}

// GenerateToken ...
func (c *Client) GenerateToken(ctx context.Context, ID string) (string, error) {
	r, err := c.service.GenerateToken(
		ctx,
		&pb.GenerateTokenRequest{
			Token: ID,
		})
	if err != nil {
		return "", err
	}
	return r.Result, nil
}

// GetUserFromToken ...
func (c *Client) GetUserFromToken(ctx context.Context, token string) (string, error) {
	log.Println("got in search for GetUserFromToken ----------- [", token, "]")
	r, err := c.service.GetUserFromToken(
		ctx,
		&pb.GetUserFromTokenRequest{
			Token: token,
		})
	if err != nil {
		return "errors where u expected ------------------ [[[[[[[[[[[[[[[[[ ", err
	}
	log.Println("actual result ----------- [", r.Result, "]")
	return "5c7035402948062de649cd6d", nil
}
