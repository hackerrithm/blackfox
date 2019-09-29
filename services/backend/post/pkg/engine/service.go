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

package engine

import (
	"context"
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/hackerrithm/blackfox/services/backend/post/pkg/domain"
)

func (s *post) Insert(ctx context.Context, author, topic, category, contentText, contentPhoto string) error {
	pst := domain.Post{}
	img := domain.Image{}
	img.Name = contentPhoto
	pst.ContentPhoto.Name = img.Name
	a := &domain.Post{
		Author:       bson.ObjectIdHex(author),
		Topic:        topic,
		Category:     category,
		ContentText:  contentText,
		ContentPhoto: pst.ContentPhoto,
	}
	if err := s.repository.Insert(ctx, *a); err != nil {
		return err
	}

	return nil
}

func (s *post) Update(ctx context.Context, author, topic, category, contentText, contentPhoto, id string) error {
	pst := domain.Post{}
	img := domain.Image{}
	img.Name = contentPhoto
	pst.ContentPhoto.Name = img.Name
	a := &domain.Post{
		Author:       bson.ObjectIdHex(author),
		Topic:        topic,
		Category:     category,
		ContentText:  contentText,
		ContentPhoto: pst.ContentPhoto,
	}
	if err := s.repository.Update(ctx, *a, id); err != nil {
		return err
	}

	return nil
}

func (s *post) FindOne(ctx context.Context, id string) (*domain.Post, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *post) ListAllPosts(ctx context.Context, skip, take uint64) (*[]domain.Post, error) {
	log.Println("got in service")
	res, err := s.repository.ListAllPosts(ctx, skip, take)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *post) Query(ctx context.Context, skip uint64, take uint64) ([]domain.Post, error) {
	// if take > 100 || (skip == 0 && take == 0) {
	// 	take = 100
	// }
	// return s.repository.List(ctx, skip, take)
	return nil, nil
}

func (s *post) Remove(ctx context.Context, id string) (string, error) {
	return s.repository.Remove(ctx, id)
}
