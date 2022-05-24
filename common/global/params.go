package global

import (
	"admin-server/common/config"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	Logger             *zap.Logger             // 日志组件
	VP                 *viper.Viper            // 配置对象
	Config             config.Server           // 配置文件
	Session            *session.Store          // 全局session池
	ConcurrencyControl = &singleflight.Group{} // 并发控制
	Db                 *gorm.DB                // 数据库Orm对象
	RedisClient        *redis.Client           // redis连接客户端
)
