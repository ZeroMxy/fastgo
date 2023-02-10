package config

import "github.com/ZeroMxy/fastgo/framework/env"

var App = map[string] string {

	// Application
	"name": env.Get("app.name", "zero"),
	"host": env.Get("app.host", "localhost:3000"),
	"mode": env.Get("app.host", "debug"),

	// Snowflake
	"worker_id": 	env.Get("snowflake.worker_id", "0"), 	// 机器 id 默认 0
	"start_time": 	env.Get("snowflake.start_time", "0"), 	// 起点时间戳 默认 0 如果在程序跑了一段时间修改了这个值 可能会导致生成相同的 id

}