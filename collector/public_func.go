package collector

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type ConfigStruct struct {
	Token  string                 `yaml:"token"`
	Port  map[string]interface{} `yaml:"port"`
	Proce map[string]interface{} `yaml:"Proce"`
}

type ParseValue struct {
	Token string                 `json:"token"`
	Port  map[string]interface{} `json:"port"`
	Proce map[string]interface{} `json:"proce"`
}

func ParseParam(parse ParseValue) {
	//var configViperConfig = viper.New()
	//configViperConfig.SetConfigName("node_exporter")
	//configViperConfig.SetConfigType("yaml")
	ParamWrite(parse.Token, parse.Port, parse.Proce)

}

func ParamWrite(token string, port, proces map[string]interface{}) {
	str := &ConfigStruct{
		Token: token,
		Port:  port,
		Proce: proces,
	}
	data, err := yaml.Marshal(str)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("node_exporter.yaml", data, 0777)
}

// 判断配置文件是否存在，不存在则创建
func IsExist() bool {
	_, err := os.Stat("./node_exporter.yaml")
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		CreateFile()
	}
	return false
}

// 创建文件函数
func CreateFile() bool {
	file, err := os.Create("./node_exporter.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
	file.Close()
	if IsExist() == true {
		return true
	}
	return false
}
