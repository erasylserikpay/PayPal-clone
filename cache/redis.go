package cache

import (
    "context"
    "github.com/go-redis/redis/v8"
    "log"
)

var ctx = context.Background()
var RedisClient *redis.Client

func Connect() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    _, err := RedisClient.Ping(ctx).Result()
    if err != nil {
        log.Fatal(err)
    }
}
