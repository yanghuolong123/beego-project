package routers

import (
	"github.com/astaxie/beego"
	"go-web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("test", &controllers.HomeController{}, "get:TestGet;post:TestPost")
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "get,post:Logout")
	//beego.AutoRouter(&controllers.LoginController{})
	//	beego.Router("/login/login", &controllers.HomeController{}, "post:Post")
}
