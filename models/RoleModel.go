package models

import "github.com/astaxie/beego/orm"

type Role struct {
	Id int64 `pk:"auto" form:"id"`
	RoleName string `orm:"size(32)" form:"roleName"`
	Remark string `orm:"size(300)" form:"remark"`
	User []*User `orm:"reverse(many)"`
	Menu []*Menu `orm:"rel(m2m)"`
}

func RoleTableName() (string){
	return "role"
}

func RolePageList(page,size int,filters map[string] interface{}) ([]*Role,int64){
	roles := make([]*Role, 0)
	qt := orm.NewOrm().QueryTable(RoleTableName())
	if len(filters) > 0 {
		for k,v := range filters{
			qt = qt.Filter(k,v)
		}
	}
	total,_:= qt.Count()
	qt.OrderBy("-id").Limit(size,(page-1)*size).All(&roles)
	return roles,total
}