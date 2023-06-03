package adapter

import (
	redis "github.com/go-redis/redis/v8"
)

const Nil = redis.Nil

type RedisClient = redis.Client

var (
	JobQueueClient *RedisClient
)

const (
	KeyJobQueueTest = "job-queue-test"
)
