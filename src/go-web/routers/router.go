package routers

import (
	"github.com/astaxie/beego"
	"go-web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home/index", &controllers.HomeController{})
}
