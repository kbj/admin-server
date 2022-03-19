package system

// CasbinRule casbin的policy存储表
type CasbinRule struct {
	Ptype  string `json:"ptype" gorm:"column:ptype"`
	RoleId string `json:"roleId" gorm:"column:v0"`
	Path   string `json:"path" gorm:"column:v1"`
	Method string `json:"method" gorm:"column:v2"`
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}
