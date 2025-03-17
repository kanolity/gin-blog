package main

import (
	"go_code/gin-vue-blog/core"
	"go_code/gin-vue-blog/global"
)

func main() {
	//读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	global.Log.Warnln("111")
	global.Log.Error("222")
	global.Log.Infof("333")
	//连接数据库
	global.DB = core.Initgorm()
}
