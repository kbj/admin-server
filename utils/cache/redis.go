// Package cache Redis连接工具类
// @author: WeiKai
// @date: 2022/5/22
package cache

import (
	"admin-server/common/global"
	"context"
	"encoding/json"
	"time"
)

// SetCacheObject 缓存基本的对象
func SetCacheObject(key string, value any, time time.Duration) error {
	// 把相关对象转为json存入redis
	object, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return global.RedisClient.Set(context.Background(), key, string(object), time).Err()
}

// GetCacheObject 获得缓存的基本对象
func GetCacheObject[T any](key string) (*T, error) {
	result, err := global.RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	// 将result转为T类型
	var param T
	err = json.Unmarshal([]byte(result), &param)
	if err != nil {
		return nil, err
	}
	return &param, nil
}

// DeleteCacheObject 删除单个对象
func DeleteCacheObject(key ...string) error {
	return global.RedisClient.Del(context.Background(), key...).Err()
}

// Expire 设置有效时间
func Expire(key string, time time.Duration) error {
	return global.RedisClient.Expire(context.Background(), key, time).Err()
}
