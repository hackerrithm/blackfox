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
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	pb "github.com/hackerrithm/blackfox/services/backend/auth/pkg/model"
)

// Client act as our model
type Client struct {
	conn    *grpc.ClientConn
	service pb.AuthServiceClient
}

// NewClient creates a new client instance
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second*13))
	if err != nil {
		return nil, fmt.Errorf("error in server !!!!!!!!!: %v", err)
	}

	c := pb.NewAuthServiceClient(conn)
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
	log.Println("got in search for GetUserFromToken -----------((here)) [", token, "]")
	defer c.conn.Close()
	r, err := c.service.GetUserFromToken(
		ctx,
		&pb.GetUserFromTokenRequest{
			Token: token,
		})
	if err != nil {
		return "", grpc.Errorf(codes.Internal, err.Error())
	}
	log.Println("actual result ----------- [", r.Result, "]")
	return r.Result, nil
}
