package base

import (
	"admin-server/common/global/constants"
	"admin-server/common/types"
	"gorm.io/gorm"
)

// Model 数据库实体通用结构
type Model struct {
	ID       *uint           `gorm:"comment:主键;primarykey" json:"id"`
	CreateAt *types.UnixTime `gorm:"comment:创建时间" json:"createAt"`
	UpdateAt *types.UnixTime `gorm:"comment:修改时间" json:"updateAt"`
	DeleteAt *gorm.DeletedAt `gorm:"comment:删除时间;index" json:"deleteAt,omitempty"`
	Sequence *int            `gorm:"comment:排序;default:0;index" json:"sequence"`
	Remark   *string         `gorm:"comment:备注;size:1000" json:"remark,omitempty"`
}

// PageModel 分页数据下的通用请求参数
type PageModel struct {
	PageNum  *int64  `json:"pageNum"`  // 要请求的页码
	PageSize *int64  `json:"pageSize"` // 每页大小
	OrderBy  *string `json:"orderBy"`  // 排序字段
}

// GetPage 获取当前页
func (p *PageModel) GetPage() int64 {
	if p.PageNum == nil {
		return constants.DefaultPage
	}
	return *p.PageNum
}

// GetPageSize 获取每页大小
func (p *PageModel) GetPageSize() int64 {
	if p.PageSize == nil {
		return constants.DefaultPageSize
	}
	return *p.PageSize
}
