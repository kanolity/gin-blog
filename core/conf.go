package core

import (
	"fmt"
	"go_code/gin-vue-blog/config"
	"go_code/gin-vue-blog/global"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

// InitConf 读取配置
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlconf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, &c)
	if err != nil {
		log.Fatalf("config Init error: %v", err)
	}
	log.Println("config yamlFile load Init success")
	log.Println(c)
	global.Config = c
}
