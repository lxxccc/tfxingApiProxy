package routers

import (
	"github.com/astaxie/beego"
	"github.com/lxxccc/tfxingApiProxy/controllers"
)

func init() {
	beego.Router("/*", &controllers.MainController{})
}
