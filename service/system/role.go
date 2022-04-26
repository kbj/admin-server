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

	// 执行分页查询
	err = p.SelectPageList(db, param.OrderBy, param.Desc)
	return err, p
}
