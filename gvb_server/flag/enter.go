package flags

import (
	"flag"
	"github.com/fatih/structs"
)

type Option struct {
	DB   bool
	User string // -u admin or -u user
}

// Parse 解析命令行参数
func Parse() Option {

	db := flag.Bool("db", false, "初始化数据库")
	user := flag.String("u", "", "创建用户")

	// 解析命令行参数写入注册的flag里
	flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

// IsWebStop 是否停止Web项目
func IsWebStop(option Option) bool {
	maps := structs.Map(&option)
	for _, val := range maps {
		switch val.(type) {
		case string:
			if val != "" {
				//f = true
				return true
			}
		case bool:
			if val != false {
				//f = true
				return true
			}
		}
	}
	//return f
	return false
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		MakeMigrations()
		return
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	flag.Usage()
}
