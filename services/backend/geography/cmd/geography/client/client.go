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
	"log"
	"time"

	"google.golang.org/grpc"

	pb "github.com/hackerrithm/blackfox/services/backend/geography/pkg/model"
)

// Client ...
type Client struct {
	conn    *grpc.ClientConn
	service pb.GeographyServiceClient
}

// NewClient ...
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second*60))
	if err != nil {
		return nil, err
	}
	c := pb.NewGeographyServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close ...
func (c *Client) Close() {
	c.conn.Close()
}

// GetLocationDistance ...
func (c *Client) GetLocationDistance(ctx context.Context, lon, lat float64) (float64, error) {
	r, err := c.service.GetLocationDistance(
		ctx,
		&pb.LocationDistanceRequest{
			Longitude: lon,
			Latitude:  lat,
		},
	)
	if err != nil {
		return 0.0, err
	}

	log.Println(r.Geography)
	return r.Geography, nil
}
