package controllers

import (
	"github.com/astaxie/beego"
	"go-web/components"
	"go-web/models"
)

type HomeController struct {
	//beego.Controller
	components.BaseController
}

func (this *HomeController) Get() {
	user := this.GetSession("user")
	this.Data["user"] = user

	//	log.Println(user)
	beego.Debug("11111111111")

	this.Layout = "layout/main.tpl"
	this.TplName = "home/index.tpl"
}

func (c *HomeController) TestGet() {
	//m := make(map[string]string)
	//m["username"] = "杨活龙"
	//m["email"] = "yhl27ml@163.com"
	//c.Data["json"] = m

	//c.ServeJSON()

	m, err := models.GetUser("1")
	if err != nil {
		c.SendRes(-1, "failed", err)
	}
	//_ = m
	c.SendRes(0, "sucess", m)
	//c.SendRes(0, "sucess", models.GetAllUser())
}

func (c *HomeController) TestPost() {
	email := c.GetString("email")
	passwd := c.GetString("password")
	c.Ctx.WriteString(email + ":" + passwd)
}
