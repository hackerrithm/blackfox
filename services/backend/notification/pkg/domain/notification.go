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

package domain

import "time"

var geographyContextKey contextKey = "geography"

type contextKey string

// Notification ...
type Notification struct {
	ID      string    `json:"id"`
	Service string    `json:"service"`
	Topic   string    `json:"topic"`
	Body    string    `json:"body"`
	Date    time.Time `json:"timestamp"`
}

// NewNotification creates a new Notification!
func NewNotification(service, topic, body string) *Notification {
	return &Notification{
		Service: service,
		Topic:   topic,
		Body:    topic,
		Date:    now(),
	}
}
