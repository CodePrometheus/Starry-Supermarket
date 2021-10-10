package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"shop-web/user-api/global"
	"time"
)

var (
	ctx = context.Background()
	min = time.Minute
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
		zap.S().Panicf("Redis连接失败失败: %s", err.Error())
	}
	return Redis
}

// Set set a capt
func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	duration := min * time.Duration(global.ServerConfig.RedisInfo.Expires)
	if err := InitRedis().Set(ctx, key, value, duration).Err(); err != nil {
		zap.S().Errorf("插入缓存失败: %s", err.Error())
	}
	return nil
}

// Get get a capt
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := InitRedis().Get(ctx, key).Result()
	if err != nil {
		zap.S().Errorf("获取缓存失败: %s", err.Error())
		return ""
	}
	if clear {
		if err := InitRedis().Del(ctx, key).Err(); err != nil {
			zap.S().Errorf("清除缓存失败: %s", err.Error())
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

func SetKey(key string, value string) {
	duration := min * time.Duration(global.ServerConfig.RedisInfo.Expires)
	if err := InitRedis().Set(ctx, key, value, duration).Err(); err != nil {
		zap.S().Errorf("插入缓存失败: %s", err.Error())
	}
}

func GetKey(key string) (string, error) {
	value, err := InitRedis().Get(ctx, key).Result()
	return value, err
}
