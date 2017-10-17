package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["username"] = "杨火龙"

	c.Layout = "layout/main.tpl"
	c.TplName = "home/index.tpl"
}

func (c *HomeController) LoginGet() {

	c.Layout = "layout/main.tpl"
	c.TplName = "home/login.tpl"
}

func (c *HomeController) LoginPost() {
}
