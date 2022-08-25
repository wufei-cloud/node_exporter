package collector

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// 存储监控参数的状态 正常true 不正常false   最终此map 返回服务端解析 是否推送告警
var MonitorisHealth = make(map[string]bool)

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
		log.Errorln(err.Error())
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
		log.Errorln(err.Error())
	}
	file.Close()
	if IsExist() == true {
		return true
	}
	return false
}

// 所有自定义监控调用入口
func AllMonitor() {
	ProcesCheck()
	PortCheck()

	fmt.Println(MonitorisHealth)

	client := &http.Client{}
	str, _ := json.Marshal(&MonitorisHealth)
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	time.AfterFunc(2*time.Second, AllMonitor)
	//request, err := http.NewRequest("POST", "http://127.0.0.1:8081/query/node_exporter", bytes.NewReader(str))
	request, err := http.NewRequest("POST", "http://127.0.0.1:8081/query/node_exporter", bytes.NewReader([]byte(str)))
	fmt.Println(request)
	request.Header.Add("Node_Token", ReadParam().GetString("token"))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//request.Header.Add("Content-Type", "application/json;charset=UTF-8")
	MonitorisHealth = make(map[string]bool)
	if err != nil {
		log.Errorln(err.Error())
	}
	response, _ := client.Do(request)
	response.Body.Close()

}
