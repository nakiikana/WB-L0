package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"tools/internals/config"
	"tools/internals/models"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"
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

func (c *Cache) OrderInfo(key uuid.UUID) (*models.Orders, error) {
	var val []byte
	fmt.Println(key)
	order := &models.Orders{}
	err := c.Connection.Get(context.Background(), key.String()).Scan(&val)
	if err == redis.Nil {
		logrus.Printf("cache: could not find the order: %v", err)
		return nil, err
	}
	err = json.Unmarshal(val, order)
	if err != nil {
		logrus.Printf("cache: could not unmarshal the order: %v", err)
		return nil, err
	}
	return order, nil
}

func (c *Cache) NewOrder(order models.Orders) error {
	newKey := order.OrderID
	orderJson, err := json.Marshal(order)
	if err != nil {
		return errors.Wrap(err, "cache: could not marshal the order")
	}
	err = c.Connection.Set(context.Background(), newKey.String(), orderJson, 0).Err()
	if err != nil {
		return errors.Wrap(err, "cache: could not save a new val to cache")
	}
	logrus.Printf("saved new value to cache: %v", order.OrderID)
	return nil
}

func (c *Cache) DataRecovery(orders []models.Orders) error {
	for _, val := range orders {
		if err := c.NewOrder(val); err != nil {
			return err
		}
	}
	return nil
}

type Order interface {
	NewOrder(order models.Orders) error
	OrderInfo(uuid uuid.UUID) (*models.Orders, error)
}
