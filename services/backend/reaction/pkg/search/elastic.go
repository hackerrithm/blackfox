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

package search

import (
	"context"
	"encoding/json"
	"log"

	"github.com/olivere/elastic"

	"github.com/hackerrithm/blackfox/services/backend/reaction/pkg/domain"
)

// ElasticRepository ...
type ElasticRepository struct {
	client *elastic.Client
}

// NewElastic ...
func NewElastic(url string) (*ElasticRepository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return &ElasticRepository{client}, nil
}

// Close ...
func (r *ElasticRepository) Close() {
}

// InsertReaction ...
func (r *ElasticRepository) InsertReaction(ctx context.Context, reaction domain.Reaction) error {
	_, err := r.client.Index().
		Index("reactions").
		Type("reaction").
		Id(reaction.ID.Hex()).
		BodyJson(reaction).
		Refresh("wait_for").
		Do(ctx)
	return err
}

// SearchReactions ...
func (r *ElasticRepository) SearchReactions(ctx context.Context, query string, skip uint64, take uint64) ([]domain.Reaction, error) {
	result, err := r.client.Search().
		Index("reactions").
		Query(
			elastic.NewMultiMatchQuery(query, "body").
				Fuzziness("3").
				PrefixLength(1).
				CutoffFrequency(0.0001),
		).
		From(int(skip)).
		Size(int(take)).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	reactions := []domain.Reaction{}
	for _, hit := range result.Hits.Hits {
		var reaction domain.Reaction
		if err = json.Unmarshal(*hit.Source, &reaction); err != nil {
			log.Println(err)
		}
		reactions = append(reactions, reaction)
	}
	return reactions, nil
}
