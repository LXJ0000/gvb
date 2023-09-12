package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/routers"
)

func main() {
	//读取配置文件
	core.InitCore()
	//初始化日志
	global.Log = core.InitLogger()
	//	连接数据库
	global.DB = core.InitGorm()

	router := routers.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server running in %s", addr)
	router.Run(addr)
}
