package components

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) SendRes(code int, msg string, data interface{}) {
	m := make(map[string]interface{})
	m["code"] = code
	m["msg"] = msg
	m["data"] = data

	this.Data["json"] = m
	this.ServeJSON()
	this.StopRun()
}
