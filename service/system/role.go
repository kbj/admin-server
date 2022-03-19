package system

import (
	"admin-server/common/global"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RoleService struct{}

// GetUserRoles 查询用户拥有的角色ID
func (r *RoleService) GetUserRoles(id int) (*string, error) {
	sql := `with recursive t as (
		select role_id from t_user_role where delete_at is null and user_id = ?
		union all
		select t1.id from t_role t1, t where t1.parent_id = t.role_id and t1.delete_at is null
	)
	select group_concat(t.role_id SEPARATOR ',') from t`
	var roleIds string
	if err := global.Db.Raw(sql, id).Scan(&roleIds).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Error("数据库执行失败！", zap.Error(err))
		return nil, err
	}
	return &roleIds, nil
}
