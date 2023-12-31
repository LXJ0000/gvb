package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"log"
	"os"
)

const ConfigFile = "settings.yaml"

// InitCore 读取yaml文件配置
func InitCore() {
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yaml Conf error: %s", err))
	}

	if err = yaml.Unmarshal(yamlConf, c); err != nil {
		log.Fatalf("config init Unmarshal %s", err)
	}
	log.Println("config yamlFile load Init success")
	global.Config = c
}

func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = os.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
