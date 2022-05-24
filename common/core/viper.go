package core

import (
	"admin-server/common/global"
	"admin-server/common/global/constants"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

// InitializeViper 初始化一个新的Viper实例
// @path 参数文件路径
func InitializeViper(path ...string) *viper.Viper {
	// 用命令行指定配置文件的路径
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(constants.ConfigEnvName); configEnv == "" {
				config = constants.ConfigFileName
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config)
			} else {
				config = configEnv
				fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	// 新建viper实例
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 添加默认值
	addDefault(v)

	// 将读取到的配置信息反序列化到配置文件单例
	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}

	// 配置文件热更新
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		// 更新配置文件单例
		if err := v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})

	return v
}

// 添加默认设置
func addDefault(v *viper.Viper) {
	v.SetDefault("zap.dir", "./logs")                                       // 日志目录
	v.SetDefault("zap.log-in-console", true)                                // 是否打印日志到控制台
	v.SetDefault("system.listen", "0.0.0.0")                                // 监听地址
	v.SetDefault("system.port", "8080")                                     // 启动端口
	v.SetDefault("system.db-type", "mysql")                                 // 默认数据库类型
	v.SetDefault("jwt.signing-key", "default-jwt-token")                    // 默认的jwt加密密钥
	v.SetDefault("jwt.expires-second", 3600)                                // 默认的token过期时间  1小时
	v.SetDefault("jwt.buffer-second", 600)                                  // 默认的token过期的缓冲时间  10分钟
	v.SetDefault("jwt.issuer", "kbj")                                       // 默认token签发者
	v.SetDefault("system.db.max-idle-conn", 10)                             // 默认数据库空闲连接数
	v.SetDefault("system.db.max-open-conn", 100)                            // 默认数据库最大连接数
	v.SetDefault("casbin.model-path", "./resources/casbin_rbac_model.conf") // casBin的原型存放位置
}
