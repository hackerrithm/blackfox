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

package mongodb

import (
	"log"

	"golang.org/x/net/context"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	cfg "github.com/hackerrithm/blackfox/services/backend/post/configs"
	"github.com/hackerrithm/blackfox/services/backend/post/pkg/domain"
	"github.com/hackerrithm/blackfox/services/backend/post/pkg/engine"
)

type (
	postRepository struct {
		session *mgo.Session
	}
)

const (
	postCollection = "post"
)

var (
	config cfg.Config
)

func newPostRepository(session *mgo.Session) engine.PostRepository {
	return &postRepository{session}
}

func (r postRepository) Insert(c context.Context, p domain.Post) error {
	s := r.session.Clone()
	defer s.Close()

	var post domain.Post

	post.Author = p.Author
	post.Category = p.Category
	post.Topic = p.Topic
	post.ContentPhoto = p.ContentPhoto
	post.ContentText = p.ContentText
	post.Agreements = p.Agreements
	post.Anonymous = p.Anonymous
	post.ContentFile = p.ContentFile
	post.Date = p.Date
	post.Followers = p.Followers
	post.Latitude = p.Latitude
	post.Longitude = p.Longitude
	post.Shares = p.Shares

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Insert(&post)
	if err != nil {
		log.Println("insert error ", err)
	}

	return nil
}

func (r postRepository) Update(c context.Context, p domain.Post, id string) error {
	s := r.session.Clone()
	defer s.Close()

	var post domain.Post

	post.Author = p.Author
	post.Category = p.Category
	post.Topic = p.Topic
	post.ContentPhoto = p.ContentPhoto
	post.ContentText = p.ContentText
	post.Agreements = p.Agreements
	post.Anonymous = p.Anonymous
	post.ContentFile = p.ContentFile
	post.Date = p.Date
	post.Followers = p.Followers
	post.Latitude = p.Latitude
	post.Longitude = p.Longitude
	post.Shares = p.Shares

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	err := col.Update(bson.M{"_id": toBSON(id)}, p)
	if err != nil {
		return err
	}

	return nil
}

func (r postRepository) Query(c context.Context, query *engine.Query) []*domain.Post {
	s := r.session.Clone()
	defer s.Close()

	col := s.DB(config.MongoDB).C(config.MongoCollection)
	p := []*domain.Post{}
	q := translateQuery(col, query)
	q.All(&p)

	return nil
}

func (r postRepository) FindOne(c context.Context, id string) (*domain.Post, error) {
	s := r.session.Clone()
	defer s.Close()

	var post *domain.Post
	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(bson.M{"_id": toBSON(id)}).One(&post)

	if err != nil {
		return nil, nil
	}

	return post, nil
}

// ListAllPosts used for finding all user posts
// by the passed skip and take parameters
func (r postRepository) ListAllPosts(ctx context.Context, skip uint64, take uint64) ([]domain.Post, error) {
	s := r.session.Clone()
	defer s.Close()
	log.Println("got in repo")

	var posts []domain.Post

	err := s.DB(config.MongoDB).C(config.MongoCollection).Find(nil).All(&posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r postRepository) Remove(c context.Context, id string) (string, error) {
	s := r.session.Clone()
	defer s.Close()

	err := s.DB(config.MongoDB).C(config.MongoCollection).Remove(bson.M{"_id": toBSON(id)}) //.One(&post)

	if err != nil {
		return "", nil
	}

	return "", nil
}
