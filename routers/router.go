package routers

import (
	"flat-goadmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/dashboard",&controllers.MainController{})

    beego.AutoRouter(&controllers.UserController{})
    beego.AutoRouter(&controllers.RoleController{})

    beego.Router("/user",&controllers.UserController{},"*:Index")
	beego.Router("/role",&controllers.RoleController{},"*:Index")
}
