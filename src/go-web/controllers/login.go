package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Data["username"] = "杨火龙"
	this.Layout = "layout/main.tpl"
	this.TplName = "login/login.tpl"
}

func (this *LoginController) Post() {
	email := this.GetString("email")
	passwd := this.GetString("password")
	//this.Ctx.WriteString(email + ":" + passwd)
	_, _ = email, passwd
	this.Ctx.Redirect(302, "/home/index")
}
