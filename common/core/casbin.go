package core

import (
	"admin-server/common/global"
	"admin-server/model/system"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"sync"
)

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func GetCasbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(global.Db)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(global.Config.Casbin.ModelPath, a)

		// 初始化默认角色
		initDefaultCasbin()
	})

	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}

// 初始化默认的通用接口权限
func initDefaultCasbin() {
	var policyCount int64
	global.Db.Model(&system.CasbinRule{}).Count(&policyCount)
	if policyCount != 0 {
		return
	}

	// 添加默认policy
	rules := [][]string{
		{"common", "/system/user/", "POST"},         // 登录接口
		{"common", "/system/menu/tree-list", "GET"}, // 菜单列表接口
	}
	ok, err := syncedEnforcer.AddPolicies(rules)
	if err != nil || !ok {
		global.Logger.Error("初始化默认接口权限失败！", zap.Error(err))
	}

	// 添加默认role
	ok, err = syncedEnforcer.AddRoleForUser("1", "common")
	if err != nil || !ok {
		global.Logger.Error("初始化默认接口权限失败！", zap.Error(err))
	}

	if err != nil || !ok {
		global.Logger.Info("初始化接口权限成功")
	} else {
		global.Logger.Info("初始化接口权限失败")
	}
}
