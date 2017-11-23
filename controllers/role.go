package controllers

import (
	m "flat-goadmin/models"
	"github.com/beego/wetalk/modules/utils"
	"github.com/astaxie/beego/orm"
)

type RoleController struct {
	BaseController
}

func (this *RoleController) Index()  {
	page,_ := this.GetInt("p",1)
	search := this.GetString("s","")
	this.Data["s"] = search
	filters := make(map[string] interface{})
	if !(search == ""){
		filters["role_name__icontains"]=search
	}
	roles,total := m.RolePageList(page,size,filters)
	paginator := utils.NewPaginator(this.Ctx.Request, size, total)
	this.Data["data"] = roles
	this.Data["paginator"] = paginator
	this.Layout = layout
	this.TplName = "role/list.html"
}

func (this *RoleController) Add()  {
	if this.IsGet(){
		this.Layout = "common/layout.html"
		this.TplName = "role/add.html"
	}else {
		role := m.Role{}
		this.ParseForm(&role)
		orm.NewOrm().Insert(&role)
		this.Redirect302("/role")
	}

}

func (this *RoleController) Del()  {
	id,_:= this.GetInt("id")
	orm.NewOrm().Delete(&m.Role{Id:id})
	this.Redirect302("/role")
}

func (this *RoleController) Edit()  {
	id,_:= this.GetInt("id")
	role := m.Role{Id:id}
	o := orm.NewOrm()
	if this.IsGet(){
		o.Read(&role)
		this.Layout = layout
		this.TplName = "role/edit.html"
		this.Data["role"] = role
	}else {
		this.ParseForm(&role)
		o.Update(&role)
		this.Redirect302("/role")
	}
}