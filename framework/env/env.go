package env

import (
	"strings"

	"gopkg.in/ini.v1"
)

type Env struct {}

// 获取 env 配置信息
func Get (name, default_value string) string {

	var file, err = ini.Load("env.ini")

	if (err != nil) {
		panic(err)
	}

	if (name == "") {
		panic("参数错误")
	}

	// 下级 name 不存在
	if (!strings.Contains(name, ".")) {
		return file.Section("").Key(name).MustString(default_value)
	}

	// 配置名称分割
	var names = strings.Split(name, ".")

	return file.Section(names[0]).Key(names[1]).MustString(default_value)

}