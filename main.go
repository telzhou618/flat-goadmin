package main

import (
	_ "flat-goadmin/routers"
	"github.com/astaxie/beego"
	"flat-goadmin/models"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"flat-goadmin/lib"
)

func main() {
	beego.AddFuncMap("timeFmt",lib.TimeFormat)
	beego.Run()
}
func init()  {
	orm.Debug = true
	models.RegModels()
	models.Connect()
	orm.RunSyncdb("default", false, false)
}

