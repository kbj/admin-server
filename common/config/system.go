package config

type System struct {
	Listen string `mapstructure:"listen" json:"listen" yaml:"listen"`   // 监听地址
	Port   string `mapstructure:"port" json:"port" yaml:"port"`         // 端口值
	DbType string `mapstructure:"db-type" json:"dbType" yaml:"db-type"` // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	Db     Db     `mapstructure:"db" json:"db" yaml:"db"`               // 数据库配置
}
