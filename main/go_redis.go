package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "sxj045115",
		DB:       0,
	})
}

func main() {
	ctx := context.Background()

	err := rdb.Set(ctx, "k1", "v1", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	result, err := rdb.Get(ctx, "k1").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
