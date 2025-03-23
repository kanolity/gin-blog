package main

import (
	"go_code/gin-vue-blog/core"
	_ "go_code/gin-vue-blog/docs"
	"go_code/gin-vue-blog/flag"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/router"
)

// @title gvb_server API文档
// @version 1.0
// @description API文档
// @host 127.0.0.1:8080
// @BasePath /

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据库
	global.DB = core.InitGorm()

	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWbeStop(option) {
		flag.SwitchOption(option)
		return
	}
	r := router.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server运行在:%s", addr)
	err := r.Run(addr)
	if err != nil {
		global.Log.Fatalf("run server error: %v", err)
	}
}
