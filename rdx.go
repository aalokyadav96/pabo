package main

import (
	"context"
	"fmt"
	"log"

	gredis "github.com/redis/go-redis/v9"
)

var conn = gredis.NewClient(&gredis.Options{
	Network: "tcp",
	Addr:    "localhost:6379",
})

func init() {
	_, err := conn.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
}

func rdxSet(key, value string) error {

	ctx := context.Background()

	_, err := conn.Set(ctx, key, value, 0).Result()
	if err != nil {
		return fmt.Errorf("error while doing SET command in gredis : %v", err)
	}

	return err

}

func rdxGet(key string) (string, error) {

	ctx := context.Background()

	value, err := conn.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("error while doing GET command in gredis : %v", err)
	}

	return value, err
}


func rdxDel(key string) (string, error) {

	ctx := context.Background()

	value, err := conn.Del(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("error while doing DEL command in gredis : %v", err)
	}

	return ""+string(value), err
}

func rdxHset(hash, key, value string) error {

	ctx := context.Background()

	_, err := conn.HSet(ctx, hash, key, value).Result()
	if err != nil {
		return fmt.Errorf("error while doing HSET command in gredis : %v", err)
	}

	return err
}

func rdxHget(hash, key string) (string) {

	ctx := context.Background()

	value, err := conn.HGet(ctx, hash, key).Result()
	if err != nil {
		return "error"
	}

	return value

}

func rdxHdel(hash, key string) (string, error) {

	ctx := context.Background()

	value, err := conn.HDel(ctx, hash, key).Result()
	if err != nil {
		return string(value), fmt.Errorf("error while doing HGET command in gredis : %v", err)
	}

	return string(value), err

}
