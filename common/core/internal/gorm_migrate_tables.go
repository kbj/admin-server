package internal

import (
	"admin-server/common/global"
	"admin-server/entity/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

// InitializeTables 自动初始化好数据库表
func InitializeTables(db *gorm.DB) {
	err := db.AutoMigrate(
		system.User{},
	)

	if err != nil {
		global.Logger.Error("初始化表结构失败！", zap.Error(err))
		os.Exit(0)
	}
	global.Logger.Info("初始化表结构成功！")
}
