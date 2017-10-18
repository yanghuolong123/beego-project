package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "go-web/routers"
)

func main() {

	var filterLogin = func(ctx *context.Context) {
		user := ctx.Input.Session("user")
		if user == nil && ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302, "/login")
		}
	}

	beego.InsertFilter("/*", beego.BeforeRouter, filterLogin)
	beego.Run()
}
