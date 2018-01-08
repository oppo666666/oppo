package routers

import (
	"github.com/astaxie/beego"
	"oppo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
