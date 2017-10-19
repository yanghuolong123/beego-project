package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	//	"github.com/astaxie/beego/logs"
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
	//log := logs.NewLogger(10000)
	//log.SetLogger(logs.AdapterFile, `{"filename":"logs/11.log"}`)
	//log.EnableFuncCallDepth(true)
	//log.Info("dss11111")
	//	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.Run()
}
