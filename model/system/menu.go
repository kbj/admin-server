package system

import "admin-server/model/common"

// Menu 菜单信息
type Menu struct {
	common.Model
	Name     *string `gorm:"comment:菜单名称;size:100;not null"`
	ParentId *uint   `gorm:"comment:上级ID;not null;default:0;index"`
	Path     *string `gorm:"comment:前端用于匹配组件的路径;size:1000;"`
	Icon     *string `gorm:"comment:图标信息;size:100"`
	IsHide   *bool   `gorm:"comment:是否隐藏;default:0"`
	Type     *uint8  `gorm:"comment:菜单类型，1是菜单2是功能;default:1;size:1"`
}
