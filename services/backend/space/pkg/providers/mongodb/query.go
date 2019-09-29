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
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/hackerrithm/blackfox/services/backend/space/pkg/engine"
)

// translateQuery converts an application query spec into
// a mongodb specific query
func translateQuery(c *mgo.Collection, query *engine.Query) *mgo.Query {
	m := bson.M{}
	for _, filter := range query.Filters {
		switch filter.Condition {
		case engine.Equal:
			m[filter.Property] = filter.Value
		case engine.LessThan:
			m[filter.Property] = bson.M{"$lt": filter.Value}
		case engine.LessThanOrEqual:
			m[filter.Property] = bson.M{"$lte": filter.Value}
		case engine.GreaterThan:
			m[filter.Property] = bson.M{"$gt": filter.Value}
		case engine.GreaterThanOrEqual:
			m[filter.Property] = bson.M{"$gte": filter.Value}
		}
	}
	q := c.Find(m)

	for _, order := range query.Orders {
		switch order.Direction {
		case engine.Ascending:
			q = q.Sort(order.Property)
		case engine.Descending:
			q = q.Sort("-" + order.Property)
		}
	}

	if query.Offset > 0 {
		q = q.Skip(query.Offset)
	}

	if query.Limit > 0 {
		q = q.Limit(query.Limit)
	}

	return q
}
