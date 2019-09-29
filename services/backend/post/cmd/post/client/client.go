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

	"github.com/hackerrithm/blackfox/services/backend/post/pkg/domain"
	pb "github.com/hackerrithm/blackfox/services/backend/post/pkg/model"
)

// Client ...
type Client struct {
	conn    *grpc.ClientConn
	service pb.PostServiceClient
}

// NewClient ...
func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithTimeout(time.Second*80))
	if err != nil {
		return nil, err
	}
	c := pb.NewPostServiceClient(conn)
	return &Client{conn, c}, nil
}

// Close ...
func (c *Client) Close() {
	c.conn.Close()
}

// Post ...
func (c *Client) Post(ctx context.Context, author, topic, category, contentText, contentPhotoName string) (string, error) {
	var pst = pb.Post{}
	var img = pb.Image{}
	img.Name = contentPhotoName
	pst.ContentPhoto = &img

	r, err := c.service.PostPost(
		ctx,
		&pb.PostPostRequest{
			Author:       author,
			Topic:        topic,
			Category:     category,
			ContentText:  contentText,
			ContentPhoto: pst.ContentPhoto,
		},
	)
	if err != nil {
		return "", err
	}

	log.Println(r.Post)
	return r.Post, nil
}

// Put ...
func (c *Client) Put(ctx context.Context, author, topic, category, contentText, contentPhotoName, id string) (string, error) {
	var pst = pb.Post{}
	var img = pb.Image{}
	img.Name = contentPhotoName
	pst.ContentPhoto = &img

	r, err := c.service.PutPost(
		ctx,
		&pb.PutPostRequest{
			Id:           id,
			Author:       author,
			Topic:        topic,
			Category:     category,
			ContentText:  contentText,
			ContentPhoto: pst.ContentPhoto,
		},
	)
	if err != nil {
		return "", err
	}
	return r.Post, nil
}

// Get ...
func (c *Client) Get(ctx context.Context, id string, userID uint64) (*domain.Post, error) {
	var pst = domain.Post{}
	var img = pb.Image{}

	r, err := c.service.GetPost(
		ctx,
		&pb.GetPostRequest{Id: id, UserID: userID},
	)
	if err != nil {
		return nil, err
	}

	img = *r.Post.ContentPhoto
	pst.ContentPhoto.Name = img.Name

	return &domain.Post{
		ID:           bson.ObjectIdHex(r.Post.Id),
		Author:       bson.ObjectIdHex(r.Post.Author),
		Topic:        r.Post.Topic,
		Category:     r.Post.Category,
		ContentText:  r.Post.ContentText,
		ContentPhoto: pst.ContentPhoto,
	}, nil
}

// GetMultiple is used to get the list of specified posts
func (c *Client) GetMultiple(ctx context.Context, skip uint64, take uint64) ([]domain.Post, error) {
	var pst = domain.Post{}
	var img = pb.Image{}
	r, err := c.service.GetMultiplePosts(
		ctx,
		&pb.GetMultiplePostsRequest{
			Skip: skip,
			Take: take,
		},
	)
	if err != nil {
		return nil, err
	}

	posts := []domain.Post{}
	for _, a := range r.Posts {
		img = *a.ContentPhoto
		pst.ContentPhoto.Name = img.Name
		posts = append(posts, domain.Post{
			ID:           bson.ObjectIdHex(a.Id),
			Author:       bson.ObjectIdHex(a.Author),
			Topic:        a.Topic,
			Category:     a.Category,
			ContentText:  a.ContentText,
			ContentPhoto: pst.ContentPhoto,
		})
	}
	return posts, nil
}

// Delete removes a post with passed identifier
func (c *Client) Delete(ctx context.Context, id string) (string, error) {
	r, err := c.service.DeletePost(
		ctx,
		&pb.DeletePostRequest{Id: id},
	)
	if err != nil {
		return "", err
	}
	return r.Id, nil
}
