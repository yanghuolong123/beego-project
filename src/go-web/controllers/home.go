package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["username"] = "杨火龙"
	c.TplName = "home/index.tpl"
}
