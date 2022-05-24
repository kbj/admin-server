package system

import (
	"admin-server/common/global"
	"admin-server/model/system"
	"admin-server/model/system/response"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MenuService struct{}

// TreeList 查询用户的菜单树
func (m *MenuService) TreeList(userId uint) (*[]response.MenuTreeModel, error) {
	// 查询出角色相关的菜单
	sql := `with recursive m as (
		select id, create_at, sequence, name, parent_id, path, icon, is_hide, type from t_menu where id in (select menu_id from t_role_menu where delete_at is null and role_id in (select role_id from t_user_role where delete_at is null and user_id = ?)) and delete_at is null
		union all
		select m1.id, m1.create_at, m1.sequence, m1.name, m1.parent_id, m1.path, m1.icon, m1.is_hide, m1.type from t_menu m1, m where m1.parent_id = m.id and m1.delete_at is null
	)
	select * from m order by sequence, create_at, id`
	var menus []system.Menu
	if err := global.Db.Raw(sql, userId).Scan(&menus).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Error("菜单树查询失败！", zap.Error(err))
		return nil, global.NewError("菜单树查询失败")
	}

	// 递归生成树结构
	return m.recursiveMenuTree(menus, 0, 1), nil
}

// 将菜单列表整理成菜单树结构
func (m MenuService) recursiveMenuTree(menus []system.Menu, parentId uint, level uint8) *[]response.MenuTreeModel {
	var treeList []response.MenuTreeModel
	for _, val := range menus {
		if *val.ParentId == parentId {
			a, err := json.Marshal(val)
			if err != nil {
				continue
			}
			var b response.MenuTreeModel
			_ = json.Unmarshal(a, &b)

			// 只有菜单类型才继续往下层搜索
			if *b.Type == 1 {
				b.Children = m.recursiveMenuTree(menus, *b.ID, level+1)
			}

			b.Level = &level
			treeList = append(treeList, b)
		}
	}
	return &treeList
}
