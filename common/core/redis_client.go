package core

import (
	"admin-server/common/global"
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"os"
)

// GetRedisClient 创建Redis客户端对象
func GetRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Error("测试连接Redis失败！程序退出", zap.Error(err))
		os.Exit(1)
		return nil
	} else {
		global.Logger.Info("测试连接Redis成功！", zap.String("pong", pong))
		return client
	}
}
