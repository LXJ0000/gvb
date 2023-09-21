package main

import (
	"gvb_server/core"
	_ "gvb_server/docs"
	flags "gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
)

// @title           GVB Server API Doc
// @version         1.0
// @description     API Doc
// @host      localhost:8080
// @BasePath  /
func main() {
	//读取配置文件
	core.InitCore()
	//初始化日志
	global.Log = core.InitLogger()
	//	连接数据库
	global.DB = core.InitGorm()
	global.Redis = core.ConnectRedis()
	//命令行参数绑定
	option := flags.Parse()
	if flags.IsWebStop(option) {
		flags.SwitchOption(option)
		return
	}
	router := routers.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server running in %s", addr)
	if err := router.Run(addr); err != nil {
		global.Log.Fatalf(err.Error())
	}

}
