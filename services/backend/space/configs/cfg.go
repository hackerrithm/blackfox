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

package configs

// Config ...
type Config struct {
	UserServiceURL    string `envconfig:"USER_SERVICE_URL"`
	ProfileServiceURL string `envconfig:"PROFILE_SERVICE_URL"`
	SpaceServiceURL   string `envconfig:"SPACE_SERVICE_URL"`
	MongoHost         string `envconfig:"MONGO_HOSTS"`
	MongoDB           string `envconfig:"MONGO_DATABASE"`
	MongoUsername     string `envconfig:"MONGO_USERNAME"`
	MongoPassword     string `envconfig:"MONGO_PASWORD"`
	MongoCollection   string `envconfig:"MONGO_COLLECTION"`
	MONGOURL          string `envconfig:"MONGO_URL"`
	GRPCPort          string `envconfig:"GRPC_PORT"`
	GRAPHQLPort       string `envconfig:"GRAPHQL_PORT"`
}
