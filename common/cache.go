package common

import (
	"github.com/go-redis/redis"
	"fmt"
	"time"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()

	if err != nil {
		fmt.Println(err.Error())
	}
}

type Cache struct {
	Key      string
	Value    string
	ExprTime time.Duration
}

func (c Cache) SetCache() {
	client.Set(c.Key, c.Value, c.ExprTime)
}

func GetCacheValueByKey(key string) string {
	value := client.Get(key).Val()

	return value
}
