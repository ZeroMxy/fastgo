package config

import "github.com/ZeroMxy/fastgo/framework/env"

var Database = map[string] string {

	// Database
	"drive":    env.Get("database.drive", "mysql"),
	"host":     env.Get("database.host", "localhost"),
	"port":     env.Get("database.port", "3306"),
	"name":     env.Get("database.name", "app"),
	"username": env.Get("database.username", "root"),
	"password": env.Get("database.password", "root"),

}