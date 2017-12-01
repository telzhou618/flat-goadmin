package routers

import (
	"flat-goadmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//首页
    beego.Router("/", &controllers.MainController{})
    beego.Router("/dashboard",&controllers.MainController{})

    //自动增删改路由
    beego.AutoRouter(&controllers.UserController{})
    beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.MenuController{})

	//分页查询路由
    beego.Router("/user",&controllers.UserController{},"*:Index")
	beego.Router("/role",&controllers.RoleController{},"*:Index")
	beego.Router("/menu",&controllers.MenuController{},"*:Index")
	beego.Router("/log",&controllers.LogController{},"*:Index")
}
