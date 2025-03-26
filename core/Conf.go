package core

import (
	"fmt"
	"go_code/gin-vue-blog/config"
	"go_code/gin-vue-blog/global"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
	"log"
)

const ConfigFile = "settings.yaml"

// InitConf 读取配置
func InitConf() {

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
	global.Config = c
}

// SetYaml 更改yaml文件
func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
