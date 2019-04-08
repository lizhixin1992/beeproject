package routers

import (
	"beeproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	//beego.AutoRouter(&controllers.MainController{})
	beego.Include(&controllers.CMSController{})
}
