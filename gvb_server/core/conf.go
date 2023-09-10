package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"io/ioutil"
	"log"
)

// InitCore 读取yaml文件配置
func InitCore() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yaml Conf error: %s", err))
	}

	if err = yaml.Unmarshal(yamlConf, c); err != nil {
		log.Fatalf("config init Unmarshal %s", err)
	}
	log.Println("config yamlFile load Init success")
	global.Config = c
}
