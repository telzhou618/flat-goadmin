package main

import (
	_ "flat-goadmin/routers"
	"github.com/astaxie/beego"
	"flat-goadmin/models"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	beego.Run()
}
func init()  {
	orm.Debug = true
	models.RegModels()
	models.Connect()
	orm.RunSyncdb("default", false, false)
}

