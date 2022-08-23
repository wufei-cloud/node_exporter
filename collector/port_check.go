package collector

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"time"
)

//  端口检查
func PortCheck() {
	for k, v := range ReadParam().GetStringMap("port") {
		cmd := exec.Command("/bin/bash", "-c", "netstat -nltpau | grep "+fmt.Sprintf("%v", v)+" | grep -v grep | grep LISTEN")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Error(k + "port  down",err.Error())
			MonitorisHealth[k] = "false" // 存在监控异常的服务存入MonitorisHealth map中 存储
			continue
		}
		MonitorisHealth[k] = "true" // 监控正常的服务存入MonitorisHealth map中 存储
	}
	time.AfterFunc(2*time.Second, PortCheck)
}
