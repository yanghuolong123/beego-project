package controllers

import (
	//	"github.com/astaxie/beego"
	"go-web/components"
	"go-web/models"
)

type LoginController struct {
	//beego.Controller
	components.BaseController
}

func (this *LoginController) Get() {
	u := this.GetSession("user")
	this.Data["user"] = u
	this.Layout = "layout/main.tpl"
	this.TplName = "login/login.tpl"
}

func (this *LoginController) Post() {
	email := this.GetString("email")
	passwd := this.GetString("password")
	//this.Ctx.WriteString(email + ":" + passwd)
	//	_, _ = email, passwd
	//	this.Ctx.Redirect(302, "/home/index")
	u, err := models.GetUserLogin(email, passwd)
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}

	this.SetSession("user", u)
	this.Ctx.Redirect(302, "/home")

	//this.SendRes(0, "login success", u)

}

func (this *LoginController) Logout() {
	this.DelSession("user")
	this.Ctx.Redirect(302, "/login")
}
