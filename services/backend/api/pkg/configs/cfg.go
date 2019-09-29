package configs

import "time"

// Config ...
type Config struct {
	UserServiceURL      string `envconfig:"USER_SERVICE_URL"`
	ProfileServiceURL   string `envconfig:"PROFILE_SERVICE_URL"`
	RedisServiceURL     string `envconfig:"REDIS_SERVICE_URL"`
	TaskServiceURL      string `envconfig:"TASK_SERVICE_URL"`
	PostServiceURL      string `envconfig:"POST_SERVICE_URL"`
	SpaceServiceURL     string `envconfig:"SPACE_SERVICE_URL"`
	GeographyServiceURL string `envconfig:"GEOGRAPHY_SERVICE_URL"`
	GoalServiceURL      string `envconfig:"GOAL_SERVICE_URL"`
	MatchServiceURL     string `envconfig:"MATCH_SERVICE_URL"`
	GroupServiceURL     string `envconfig:"GROUP_SERVICE_URL"`
	Server              struct {
		Port         string        `envconfig:"SERVER_PORT"`
		ReadTimeout  time.Duration `envconfig:"READ_TIMEOUT"`
		WriteTimeout time.Duration `envconfig:"WRITE_TIMEOUT"`
		IdleTimeout  time.Duration `envconfig:"IDLE_TIMEOUT"`
	}
}
