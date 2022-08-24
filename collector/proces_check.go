package collector

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os/exec"
)

// 从yaml配置文件读取配置
func ReadParam() *viper.Viper {
	var configViperConfig = viper.New()
	configViperConfig.AddConfigPath("./")
	configViperConfig.SetConfigName("node_exporter")
	configViperConfig.SetConfigType("yaml")
	if err := configViperConfig.ReadInConfig(); err != nil {
		log.Errorln(err.Error())
	}
	return configViperConfig
}

// 从配置文件获取进程信息  监控进程是否存在
func ProcesCheck() {
	MonitorisHealth := make(map[string]string)
	for k, v := range ReadParam().GetStringMap("Proce") {
		cmd := exec.Command("/bin/bash", "-c", "ps -ef | grep "+fmt.Sprintf("%v", v)+" | grep -v grep")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Errorln(k+" proce down ", err.Error())
			MonitorisHealth[k] = "false" // 存在监控异常的服务存入MonitorisHealth map中 存储
			continue
		}
		MonitorisHealth[k] = "true" // 监控正常的服务存入MonitorisHealth map中 存储
	}
}
