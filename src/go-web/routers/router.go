package routers

import (
	"github.com/astaxie/beego"
	"go-web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home/index", &controllers.HomeController{})
	beego.Router("/home/login", &controllers.HomeController{}, "get:LoginGet;post:LoginPost")
	//	beego.Router("/home/login", &controllers.HomeController{}, "post:LoginPost")
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "get,post:Logout")
	//beego.AutoRouter(&controllers.LoginController{})
	//	beego.Router("/login/login", &controllers.HomeController{}, "post:Post")
}
