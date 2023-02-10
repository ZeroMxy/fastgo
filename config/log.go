package config

import "github.com/ZeroMxy/fastgo/framework/env"

var Log = map[string]string{

	// Log
	"path":   env.Get("log.path", "storage/log/"), // 日志保存路径
	"model":  env.Get("log.model", "single"),      // single 单文件 daily 按日期切割文件
	"day":    env.Get("log.day", "7"),             // 日志保留天数
	"system": env.Get("system", "false"),          // api 请求日志 true 开启 false 关闭

}