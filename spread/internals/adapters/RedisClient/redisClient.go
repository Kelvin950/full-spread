package redisclient

import (
	"github.com/redis/go-redis/v9"
)

type IRedisClient interface {
}

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(addr, password string) *RedisClient {

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})

	return &RedisClient{
		client: client,
	}

}

func (r RedisClient) MakeRedisClient() interface{} {
	return r.client
}
