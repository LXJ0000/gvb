package main

import (
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	//读取配置文件
	core.InitCore()
	//	连接数据库
	global.DB = core.InitGorm()
}
