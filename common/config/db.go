package config

type Db struct {
	Host        string `mapstructure:"host" json:"host" yaml:"host"`                          // 地址
	Port        string `mapstructure:"port" json:"port" yaml:"port"`                          // 端口
	DbName      string `mapstructure:"db-name" json:"dbName" yaml:"db-name"`                  // 数据库名
	UserName    string `mapstructure:"user-name" json:"userName" yaml:"user-name"`            // 用户名
	Password    string `mapstructure:"password" json:"password" yaml:"password"`              // 密码
	Config      string `mapstructure:"config" json:"config" yaml:"config"`                    // 高级配置
	MaxIdleConn int    `mapstructure:"max-idle-conn" json:"maxIdleConn" yaml:"max-idle-conn"` // 空闲中的最大连接数
	MaxOpenConn int    `mapstructure:"max-open-conn" json:"maxOpenConn" yaml:"max-open-conn"` // 打开到数据库的最大连接数
}
