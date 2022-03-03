package core

import (
	"admin-server/common/core/internal"
	"admin-server/common/global"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
	"strings"
)

// InitializeDbInstance 初始化数据库对象
func InitializeDbInstance() *gorm.DB {
	var db *gorm.DB
	switch strings.ToLower(global.Config.System.DbType) {
	case "mysql":
		db = initializeMySQLInstance()
		break
	case "pgsql":
		db = initializeMySQLInstance()
		break
	default:
		db = initializeMySQLInstance()
		break
	}

	// 自动迁移表结构
	internal.InitializeTables(db)

	return db
}

// 生成gorm的配置信息
func getGormConfig() *gorm.Config {
	// 禁用自动创建外键约束
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: false}

	// 修改默认的命名策略
	config.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   "t_", // 表名前缀
		SingularTable: true, // 使用单数表名
	}

	// 设置日志
	log := zapgorm2.New(global.Logger)
	log.SetAsDefault()
	config.Logger = log

	return config
}
