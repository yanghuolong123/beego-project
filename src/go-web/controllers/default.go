package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["username"] = "杨火龙"
	c.TplName = "default/index.tpl"
}
