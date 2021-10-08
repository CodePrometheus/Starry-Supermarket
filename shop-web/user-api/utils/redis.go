package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"shop-web/user-api/global"
	"time"
)

var (
	ctx = context.Background()
)

const CAPTCHA = "captcha:"

type RedisStore struct {
}

func InitRedis() *redis.Client {
	config := global.ServerConfig.RedisInfo
	Redis := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})
	if err := Redis.Ping(ctx).Err(); err != nil {
		zap.S().Errorw("Redis连接失败失败: ", err.Error())
	}
	return Redis
}

// Set set a capt
func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	fmt.Println("key: " + key)
	fmt.Println("value: " + value)
	if err := InitRedis().Set(ctx, key, value, time.Minute*2).Err(); err != nil {
		zap.S().Errorw("插入缓存失败: ", err.Error())
	}
	return nil
}

// Get get a capt
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := InitRedis().Get(ctx, key).Result()
	if err != nil {
		zap.S().Errorw("获取缓存失败: ", err.Error())
		return ""
	}
	if clear {
		if err := InitRedis().Del(ctx, key).Err(); err != nil {
			zap.S().Errorw("清除缓存失败: ", err.Error())
			return ""
		}
	}
	return val
}

// Verify verify a capt
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
