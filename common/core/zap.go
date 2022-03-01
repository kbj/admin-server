package core

import (
	"admin-server/common/global"
	"admin-server/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// InitializeZap 初始化zap实例
func InitializeZap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.Config.Zap.Dir); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.Config.Zap.Dir)
		_ = os.Mkdir(global.Config.Zap.Dir, os.ModePerm)
	}

	// 创建各个级别，并把error以上更高级的合并到error中
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	// 依据不同级别生成不同的core
	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("%s/server_debug.log", global.Config.Zap.Dir), debugPriority),
		getEncoderCore(fmt.Sprintf("%s/server_info.log", global.Config.Zap.Dir), infoPriority),
		getEncoderCore(fmt.Sprintf("%s/server_warn.log", global.Config.Zap.Dir), warnPriority),
		getEncoderCore(fmt.Sprintf("%s/server_error.log", global.Config.Zap.Dir), errorPriority),
	}

	// 生成zap对象
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
	return logger
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(logPath string, level zapcore.LevelEnabler) zapcore.Core {
	// 使用lumberjack切分log
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logPath, // 日志文件的位置
		MaxSize:    10,      // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 200,     // 保留旧文件的最大个数
		MaxAge:     90,      // 保留旧文件的最大天数
		Compress:   false,   // 是否压缩/归档旧文件
	}

	// 读取配置文件，是否要打印信息到控制台
	var writer zapcore.WriteSyncer
	if global.Config.Zap.LogInConsole {
		// 同时打印到控制台和文件，组成Writer
		writer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	} else {
		writer = zapcore.AddSync(lumberJackLogger)
	}

	return zapcore.NewCore(getEncoderConfig(), writer, level)
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	//switch {
	//case global.GVA_CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
	//	config.EncodeLevel = zapcore.LowercaseLevelEncoder
	//case global.GVA_CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
	//	config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	//case global.GVA_CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
	//	config.EncodeLevel = zapcore.CapitalLevelEncoder
	//case global.GVA_CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
	//	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//default:
	//	config.EncodeLevel = zapcore.LowercaseLevelEncoder
	//}

	// 现在先使用控制台的编码方式，还有JSON编码器:NewJSONEncoder
	encoder := zapcore.NewConsoleEncoder(config)
	return encoder
}
