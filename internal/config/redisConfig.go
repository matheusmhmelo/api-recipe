package config

import "time"

type RedisConfig struct {
	Host		string	`yaml:"host"`
	Port		string 	`yaml:"port"`
	Password	string	`yaml:"password"`
	Database	string	`yaml:"database"`
	Duration    string	`yaml:"duration"`
}

const DefaultTimeDuration time.Duration = 1800000000000

func (redis *RedisConfig) GetTimeDuration() time.Duration {
	duration, err := time.ParseDuration(redis.Duration)
	if err != nil {
		duration = DefaultTimeDuration
	}
	return duration
}