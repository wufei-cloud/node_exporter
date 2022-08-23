package collector

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// 存储监控参数的状态 正常true 不正常false   最终此map 返回服务端解析 是否推送告警
var MonitorisHealth = make(map[string]string)

// 写入yaml 配置文件结构体
type ConfigStruct struct {
	Token string            `yaml:"token"`
	Port  map[string]string `yaml:"port"`
	Proce map[string]string `yaml:"Proce"`
}

// 解析json 类型结构体
type ParseValue struct {
	Token string            `json:"token"`
	Port  map[string]string `json:"port"`
	Proce map[string]string `json:"proce"`
}

// 接收服务端传入的参数 解析后写入yaml 配置文件
func ParseParam(parse ParseValue) {
	ParamWrite(parse.Token, parse.Port, parse.Proce)

}

// 将解析到的参数写入到yaml 配置文件
func ParamWrite(token string, port, proces map[string]string) {
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
