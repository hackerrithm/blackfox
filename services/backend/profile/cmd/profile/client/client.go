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
	"gopkg.in/mgo.v2/bson"

	"github.com/hackerrithm/blackfox/services/backend/profile/pkg/domain"
	pb "github.com/hackerrithm/blackfox/services/backend/profile/pkg/model"
)

// Client ...
type Client struct {
	conn    *grpc.ClientConn
	service pb.ProfileServiceClient
}

// NewClient ...
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second*100))
	if err != nil {
		return nil, err
	}
	c := pb.NewProfileServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close ...
func (c *Client) Close() {
	c.conn.Close()
}

// Post ...
func (c *Client) Post(ctx context.Context, username, level, about string, followers, following []string, rings int32) (string, error) {
	r, err := c.service.Post(
		ctx,
		&pb.PostRequest{
			Username:  username,
			About:     about,
			Rings:     rings,
			Level:     level,
			Followers: followers,
			Following: following,
		},
	)
	if err != nil {
		return "", err
	}

	log.Println(r.Profile)
	return r.Profile, nil
}

// Put ...
func (c *Client) Put(ctx context.Context, id, username, level, about, profileImage, backgroundImage string, followers, following []string, rings int32) (string, error) {
	var pst = pb.Profile{}
	var bkgimg = pb.Image{}
	var profimg = pb.Image{}
	profimg.Name = profileImage
	bkgimg.Name = backgroundImage
	pst.ProfileImg = &profimg
	pst.BackgroundImg = &bkgimg

	r, err := c.service.Put(
		ctx,
		&pb.PutRequest{
			Id:            id,
			Username:      username,
			About:         about,
			Rings:         rings,
			Level:         level,
			Followers:     followers,
			Following:     following,
			BackgroundImg: pst.BackgroundImg,
			ProfileImg:    pst.ProfileImg,
		},
	)
	if err != nil {
		return "", err
	}
	return r.Profile, nil
}

// Get ...
func (c *Client) Get(ctx context.Context, id string, userID uint64) (*domain.Profile, error) {
	var prof = domain.Profile{}
	var profimg = pb.Image{}
	var bkgimg = pb.Image{}

	r, err := c.service.Get(
		ctx,
		&pb.GetRequest{Id: id, UserID: userID},
	)
	if err != nil {
		return nil, err
	}

	bkgimg = *r.Profile.BackgroundImg
	profimg = *r.Profile.ProfileImg
	prof.ProfileImage.Name = profimg.Name
	prof.BackgroundImage.Name = bkgimg.Name

	var fllwrs, fllwing []bson.ObjectId
	for _, f := range r.Profile.Followers {
		fllwrs = append(fllwrs, bson.ObjectIdHex(f))
	}

	for _, f := range r.Profile.Following {
		fllwing = append(fllwing, bson.ObjectIdHex(f))
	}

	return &domain.Profile{
		ID:        bson.ObjectIdHex(r.Profile.Id),
		UserName:  r.Profile.Username,
		About:     r.Profile.About,
		Level:     r.Profile.Level,
		Rings:     r.Profile.Rings,
		Followers: fllwrs,
		Following: fllwing,
	}, nil
}

// GetMultiple is used to get the list of specified profiles
func (c *Client) GetMultiple(ctx context.Context, skip uint64, take uint64) ([]domain.Profile, error) {
	var fllwrs, fllwing []bson.ObjectId
	var prof = domain.Profile{}
	var profimg = pb.Image{}
	var bkgimg = pb.Image{}

	r, err := c.service.GetMultiple(
		ctx,
		&pb.GetMultipleRequest{
			Skip: skip,
			Take: take,
		},
	)
	if err != nil {
		return nil, err
	}

	profiles := []domain.Profile{}
	for _, a := range r.Profiles {
		bkgimg = *a.BackgroundImg
		profimg = *a.ProfileImg
		prof.ProfileImage.Name = profimg.Name
		prof.BackgroundImage.Name = bkgimg.Name

		for _, f := range a.Followers {
			fllwrs = append(fllwrs, bson.ObjectIdHex(f))
		}

		for _, f := range a.Following {
			fllwing = append(fllwing, bson.ObjectIdHex(f))
		}

		profiles = append(profiles, domain.Profile{
			ID:              bson.ObjectIdHex(a.Id),
			UserName:        a.Username,
			About:           a.About,
			Level:           a.Level,
			Rings:           a.Rings,
			Followers:       fllwrs,
			Following:       fllwing,
			ProfileImage:    prof.ProfileImage,
			BackgroundImage: prof.BackgroundImage,
		})
	}
	return profiles, nil
}

// Delete removes a profile with passed identifier
func (c *Client) Delete(ctx context.Context, id string) (string, error) {
	r, err := c.service.Delete(
		ctx,
		&pb.DeleteRequest{Id: id},
	)
	if err != nil {
		return "", err
	}
	return r.Id, nil
}
