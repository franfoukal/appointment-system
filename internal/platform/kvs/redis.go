package kvs

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	customerror "github.com/labscool/mb-appointment-system/internal/feature/custom"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
)

type (
	RedisClient struct {
		client *redis.Client
	}

	clientOptions struct {
		password string
		db       int
	}

	clientOption interface {
		applyOption(opts *clientOptions)
	}

	optFunc func(opts *clientOptions)
)

var (
	configDefault = clientOptions{
		password: "",
		db:       0,
	}
)

func NewClient(addr string, opts ...clientOption) (*RedisClient, error) {
	config := configDefault // copy
	for _, opt := range opts {
		opt.applyOption(&config)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.password,
		DB:       config.db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("unable to ping to redis server")
	}

	return &RedisClient{
		client,
	}, nil
}

func (c *RedisClient) Get(ctx context.Context, key string) (interface{}, error) {
	val, err := c.client.Get(key).Result()

	switch {
	case err == redis.Nil:
		return nil, customerror.EntityNotFoundError("")
	case err != nil:
		logger.Errorf("error getting item from kvs: %s", err.Error())
		return nil, fmt.Errorf("error getting item from kvs")
	}

	return val, nil
}

func (c *RedisClient) MGet(keys []string) ([]interface{}, error) {
	return c.client.MGet(keys...).Result()
}

func (c *RedisClient) Set(key string, value interface{}, expiration_min int64) error {
	exp := time.Duration(time.Duration(expiration_min) * time.Minute)
	return c.client.Set(key, value, exp).Err()
}

func (c *RedisClient) MSet(keys []string, values []interface{}, expiration_min int64) error {
	exp := time.Duration(time.Duration(expiration_min) * time.Second)
	var ifaces []interface{}
	pipe := c.client.TxPipeline()
	for i := range keys {
		ifaces = append(ifaces, keys[i], values[i])
		pipe.Expire(keys[i], exp)
	}

	if err := c.client.MSet(ifaces...).Err(); err != nil {
		return err
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func (f optFunc) applyOption(o *clientOptions) {
	f(o)
}

func WithPassword(password string) clientOption {
	return optFunc(func(options *clientOptions) {
		options.password = password
	})
}

func WithDB(db int) clientOption {
	return optFunc(func(options *clientOptions) {
		options.db = db
	})
}
