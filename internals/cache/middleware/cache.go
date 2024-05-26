package cache

import (
	"context"
	"fmt"
	"tools/internals/config"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type Cache struct {
	Connection *redis.Client
}

func NewCache(config *config.Configuration) *Cache {
	connection := NewRedisConnection(config)
	return &Cache{Connection: connection}
}

func NewRedisConnection(config *config.Configuration) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       0,
	})
	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		logrus.Errorf("Error when pinging redis: %v", err)
		return nil
	}
	logrus.Printf("Successfully started a new redis connection: %s\n", ping)
	return client
}

// func (c *Cache) NewOrder(order models.Orders) error {

// }
