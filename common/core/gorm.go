package core

import (
	"admin-server/common/core/internal"
	"admin-server/common/global"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
	"strings"
)

// InitializeDbInstance 初始化数据库对象
func InitializeDbInstance() *gorm.DB {
	var db *gorm.DB
	switch strings.ToLower(global.Config.System.DbType) {
	case "mysql":
		db = initializeMySQLInstance()
		break
	case "pgsql":
		db = initializeMySQLInstance()
		break
	default:
		db = initializeMySQLInstance()
		break
	}

	// 自动迁移表结构
	internal.InitializeTables(db)

	return db
}

// 生成gorm的配置信息
func getGormConfig() *gorm.Config {
	// 禁用自动创建外键约束
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: false}

	// 修改默认的命名策略
	config.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   "t_", // 表名前缀
		SingularTable: true, // 使用单数表名
	}

	// 设置日志
	log := zapgorm2.New(global.Logger)
	log.SetAsDefault()
	config.Logger = log

	return config
}

// PageResult 分页数据结构
type PageResult[T any] struct {
	Records  []T   `json:"records"`  // 分页数据
	Total    int64 `json:"total"`    // 总条目数
	Pages    int64 `json:"pages"`    // 总页数
	PageSize int64 `json:"pageSize"` // 每页数据量
	Current  int64 `json:"current"`  // 当前页数
}

// SelectPageList 分页查询方法
func (p *PageResult[T]) SelectPageList(db *gorm.DB) error {
	db.Count(&p.Total)
	if p.Total == 0 {
		// 没有符合条件的数据，直接返回一个T类型的空列表
		p.Records = []T{}
	} else {
		if err := db.Scopes(paginate(p)).Find(&p.Records).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有符合条件的数据，直接返回一个T类型的空列表
			p.Records = []T{}
		} else {
			return err
		}
	}
	return nil
}

// Paginate 泛型封装的分页抽象
func paginate[T any](page *PageResult[T]) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page.Current <= 0 {
			page.Current = 0
		}
		switch {
		case page.PageSize > 100:
			page.PageSize = 100
		case page.PageSize <= 0:
			page.PageSize = 10
		}
		page.Pages = page.Total / page.PageSize
		if page.Total%page.PageSize != 0 {
			page.Pages++
		}
		p := page.Current
		if page.Current > page.Pages {
			p = page.Pages
		}
		size := page.PageSize
		offset := int((p - 1) * size)
		return db.Offset(offset).Limit(int(size))
	}
}
