package config

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 连接密码
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // 连接的数据库序号
}
