package models

import (
	"github.com/astaxie/beego/orm"
)

type Menu struct {
	Id int64	`pk:"auto"`
	Text string `orm:"size(30)"`
	Icon string `orm:"size(100)"`
	Code string `orm:"size(20)"`
	Uri string `orm:"size(300)"`
	Resource string `orm:size(50)`
	Pid int64
}

func MenuTableName() (string){
	return "menu"
}
func MenuPageList(page,size int,filters map[string] interface{})  ([]*Menu,int64){
	menus := make([]*Menu, 0)
	qt := orm.NewOrm().QueryTable(MenuTableName())
	if len(filters) > 0 {
		for k,v := range filters{
			qt = qt.Filter(k,v)
		}
	}
	total,_:= qt.Count()
	qt.OrderBy("code").Limit(size,(page-1)*size).All(&menus)
	return menus,total
}
