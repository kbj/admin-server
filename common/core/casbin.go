package core

import (
	"admin-server/common/global"
	"admin-server/model/system"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"strings"
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
		syncedEnforcer.AddFunction("check_roles", checkRolesFunc)

		// 初始化默认角色
		initDefaultCasbin()
	})

	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}

// 把传入的资源用逗号连接的角色ID都分开匹配
func checkRolesFunc(args ...interface{}) (interface{}, error) {
	r := args[0].(string)
	p := args[1].(string)

	// 判断p是*的话直接通过
	if "*" == p {
		return true, nil
	}

	// 把r用逗号分割
	roles := strings.Split(r, ",")
	for i := range roles {
		if roles[i] == p {
			return true, nil
		}
	}

	return false, nil
}

// 初始化默认的通用接口权限
func initDefaultCasbin() {
	var policyCount int64
	global.Db.Model(&system.CasbinRule{}).Count(&policyCount)
	if policyCount != 0 {
		return
	}
	rules := [][]string{
		{"*", "/system/user/", "POST"},         // 登录接口
		{"*", "/system/menu/tree_list", "GET"}, // 菜单列表接口
	}
	ok, err := syncedEnforcer.AddPolicies(rules)
	if err != nil {
		global.Logger.Error("初始化默认接口权限失败！", zap.Error(err))
	}
	if ok {
		global.Logger.Info("初始化接口权限成功")
	} else {
		global.Logger.Info("初始化接口权限失败")
	}
}
