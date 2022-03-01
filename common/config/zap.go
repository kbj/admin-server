package config

type Zap struct {
	Dir          string `mapstructure:"dir" json:"dir"  yaml:"dir"`                               // 日志文件夹
	LogInConsole bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"` // 输出控制台
}
