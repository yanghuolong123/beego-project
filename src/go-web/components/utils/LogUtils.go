package utils

import (
	"github.com/astaxie/beego/logs"
)

func Logs() (log *logs.BeeLogger) {
	log = logs.NewLogger(10000)
	log.SetLogger(logs.AdapterFile, `{"filename":"/var/log/go-web/go-web.log"}`)
	log.EnableFuncCallDepth(true)

	return
}
