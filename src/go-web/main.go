package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	beeUtils "github.com/astaxie/beego/utils"
	_ "github.com/go-sql-driver/mysql"
	_ "go-web/routers"
)

func init() {
	dbhost := beego.AppConfig.String("mysql.host")
	dbport := beego.AppConfig.String("mysql.port")
	dbuser := beego.AppConfig.String("mysql.user")
	dbpasswd := beego.AppConfig.String("mysql.pass")
	dbname := beego.AppConfig.String("mysql.dbname")

	conn := dbuser + ":" + dbpasswd + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", conn)
}

func main() {
	filterRoute := []string{"/login", "/test", "/webupload"}
	var filterLogin = func(ctx *context.Context) {
		user := ctx.Input.Session("user")
		if user == nil && !beeUtils.InSlice(ctx.Request.RequestURI, filterRoute) {
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
