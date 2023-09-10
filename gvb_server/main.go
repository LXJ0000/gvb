package main

import (
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	//读取配置文件
	core.InitCore()
	//初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("hhhh")
	global.Log.Error("hhhh")
	global.Log.Infof("hhhh")
	//	连接数据库
	global.DB = core.InitGorm()
}
