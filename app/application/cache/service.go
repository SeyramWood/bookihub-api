package cache

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	rediscache "github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"

	"github.com/SeyramWood/bookibus/app/adapters/gateways"
	"github.com/SeyramWood/bookibus/config"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

type cache struct {
	client *rediscache.Cache
}

func NewService() gateways.CacheService {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:6379", config.App().Redis),
			Password: "", // no password set
			DB:       0,  // use default DB
		},
	)
	_, e := rdb.Ping(context.Background()).Result()
	if e != nil {
		log.Println(e)
	}
	mycache := rediscache.New(
		&rediscache.Options{
			Redis:      rdb,
			LocalCache: rediscache.NewTinyLFU(1000, time.Minute),
		},
	)
	return &cache{client: mycache}
}

func (c *cache) Set(key string, value any, ttl time.Duration) error {
	if err := c.client.Set(
		&rediscache.Item{
			Ctx:   context.Background(),
			Key:   key,
			Value: value,
			TTL:   ttl,
		},
	); err != nil {
		return err
	}
	return nil
}
func (c *cache) Get(key string, obj any) error {
	if err := c.client.Get(context.Background(), key, obj); err != nil {
		return ErrKeyNotFound
	}
	return nil
}
func (c *cache) Exist(key string) bool {
	return c.client.Exists(context.Background(), key)
}
func (c *cache) Delete(key string) error {
	return c.client.Delete(context.Background(), key)
}
