package system

import (
	"admin-server/common/core"
	"admin-server/common/global"
	"admin-server/model/system"
	"admin-server/model/system/request"
	"admin-server/model/system/response"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RoleService struct{}

// GetUserRoles 查询用户拥有的角色ID
func (r *RoleService) GetUserRoles(id int) (*[]int, error) {
	sql := `select t.role_id from t_user_role t where t.delete_at is null and t.user_id = ?`
	var roleIds []int
	if err := global.Db.Raw(sql, id).Scan(&roleIds).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Error("数据库执行失败！", zap.Error(err))
		return nil, err
	}
	return &roleIds, nil
}

// GetRoleList 查询角色列表
func (r *RoleService) GetRoleList(param *request.SysRoleParamModel) (err error, result *core.PageResult[response.SysRoleResponse]) {
	// 构建分页对象
	p := &core.PageResult[response.SysRoleResponse]{
		Current:  param.GetPage(),
		PageSize: param.GetPageSize(),
	}

	// 构建查询对象
	db := global.Db.Model(&system.Role{})
	if param.RoleName != nil && *param.RoleName != "" {
		db.Where("role_name like ?", "%"+*param.RoleName+"%")
	}
	if param.RoleCode != nil && *param.RoleCode != "" {
		db.Where("role_code like ?", "%"+*param.RoleCode+"%")
	}
	if param.CreateAt != nil && param.CreateAt[0] != nil && param.CreateAt[1] != nil {
		db.Where("create_at >= ? and create_at <= ?", param.CreateAt[0].ToTime(), param.CreateAt[1].ToTime())
	}

	// 执行分页查询
	err = p.SelectPageList(db, param.OrderBy, param.Desc)
	return err, p
}

// Delete 删除角色
func (r *RoleService) Delete(id *[]uint) error {
	// 检查角色是否已经有用户
	var roles []system.Role
	if err := global.Db.Where("id in ?", *id).Find(&roles).Error; err != nil {
		return err
	}
	e := core.GetCasbin()
	for _, role := range roles {
		users, err := e.GetUsersForRole(*role.RoleCode)
		if (users != nil && len(users) > 0) || err != nil {
			return global.NewError(*role.RoleName + "角色下已关联有用户，无法删除")
		}
	}

	return global.Db.Transaction(func(tx *gorm.DB) error {
		// 先删除角色菜单关联
		for _, role := range roles {
			_, err := e.DeleteRole(*role.RoleCode)
			if err != nil {
				return err
			}
		}

		// 删除角色
		if err := tx.Delete(&system.Role{}, *id).Error; err != nil {
			return err
		}
		return nil
	})
}
