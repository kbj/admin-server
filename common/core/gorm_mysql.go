package core

import (
	"admin-server/common/global"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// 初始化MySQL的数据库连接对象
func initializeMySQLInstance() *gorm.DB {
	if global.Config.System.Db.DbName == "" || global.Config.System.Db.Password == "" {
		global.Logger.Error("MySQL数据库连接信息未配置，退出程序")
		os.Exit(0)
	}

	mysqlConfig := mysql.Config{
		DSN:                       getMySQLDsn(),
		Conn:                      nil,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         500,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), getGormConfig()); err != nil {
		global.Logger.Panic("创建数据库连接失败！", zap.Error(err))
		os.Exit(0)
		return nil
	} else {
		// 设置连接池信息
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(global.Config.System.Db.MaxIdleConn)
		sqlDB.SetMaxOpenConns(global.Config.System.Db.MaxOpenConn)
		return db
	}
}

// 获取MySQL的连接地址
func getMySQLDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		global.Config.System.Db.UserName,
		global.Config.System.Db.Password,
		global.Config.System.Db.Host,
		global.Config.System.Db.Port,
		global.Config.System.Db.DbName,
		global.Config.System.Db.Config,
	)
}
