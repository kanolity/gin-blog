package main

import (
	"fmt"
	"go_code/gin-vue-blog/core"
	"go_code/gin-vue-blog/global"
)

func main() {
	//读取配置文件
	core.InitConf()
	fmt.Println(global.Config)
}
