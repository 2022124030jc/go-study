package main

// 测试连接 Redis 数据库
// Redis 是一个开源的内存数据结构存储系统，可以用作数据库、缓存和消息代理
// https://golang.halfiisland.com/community/database/Redis.html

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestQuickStart(t *testing.T) {
	ctx := context.Background()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer redisClient.Close()

	// 设置键值对
	err := redisClient.Set(ctx, "hello", "world", 0).Err()
	if err != nil {
		t.Fatal(err)
	}

	// 读取值
	result, err := redisClient.Get(ctx, "hello").Result()
	if err == redis.Nil {
		t.Log("key not exist")
	} else if err != nil {
		t.Fatal(err)
	}
	t.Log("result:", result)
}
