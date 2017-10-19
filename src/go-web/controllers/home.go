package controllers

import (
	"github.com/astaxie/beego"
	"go-web/components/ext"
	"go-web/components/utils"
	"go-web/models"
)

type HomeController struct {
	//beego.Controller
	ext.BaseController
}

func (this *HomeController) Get() {
	user := this.GetSession("user")
	this.Data["user"] = user

	utils.Logs().Info("1111weeee")

	this.Layout = "layout/main.tpl"
	this.TplName = "home/index.tpl"
}

func (c *HomeController) TestGet() {
	//m := make(map[string]string)
	//m["username"] = "杨活龙"
	//m["email"] = "yhl27ml@163.com"
	//c.Data["json"] = m

	//c.ServeJSON()

	dbhost := beego.AppConfig.String("mysqlhost")
	utils.Logs().Info("==== mysqlhost:" + dbhost)
	utils.Logs().Info(dbhost)

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
