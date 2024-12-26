package services

import (
    "context"
    "encoding/json"
    "paypal-clone/cache"
    "time"
)

var ctx = context.Background()

func SetCache(key string, value interface{}, duration time.Duration) error {
    jsonData, err := json.Marshal(value)
    if err != nil {
        return err
    }
    return cache.RedisClient.Set(ctx, key, jsonData, duration).Err()
}

func GetCache(key string, dest interface{}) error {
    jsonData, err := cache.RedisClient.Get(ctx, key).Result()
    if err != nil {
        return err
    }
    return json.Unmarshal([]byte(jsonData), dest)
}
