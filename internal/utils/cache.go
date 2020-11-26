package utils

import (
	"errors"
	"github.com/matheusmhmelo/api-recipe/internal/config"
	redisClient "gopkg.in/redis.v5"
	"log"
	"strconv"
	"time"
)

type Redis struct {
	Client *redisClient.Client
	Db     config.RedisConfig
}

func NewRedis(config config.RedisConfig) (*Redis, error) {
	red := Redis{Db: config}
	err := red.Connect()
	if err != nil {
		return nil, err
	}
	return &red, nil
}

func (r *Redis) Connect() error {
	db, _ := strconv.Atoi(r.Db.Database)
	r.Client = redisClient.NewClient(&redisClient.Options{
		Addr:     r.Db.Host + ":" + r.Db.Port,
		Password: r.Db.Password,
		DB:       db,
	})
	_, err := r.Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Set(key, value string, duration time.Duration) error {
	_, err := r.Client.Set(key, value, duration).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Get(key string) (string, error) {
	value, err := r.Client.Get(key).Result()
	if err == redisClient.Nil {
		return "", errors.New("key not found")
	} else if err != nil {
		return "", err
	}
	return value, nil
}

func (r *Redis) Close() {
	err := r.Client.Close()
	if err != nil {
		log.Println(err)
	}
}