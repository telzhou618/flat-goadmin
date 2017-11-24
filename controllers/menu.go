package controllers

import (
	m "flat-goadmin/models"
	"github.com/beego/wetalk/modules/utils"
)

type MenuController struct {
	BaseController
}

func (this *MenuController) Index()  {
	page,_ := this.GetInt("p",1)
	search := this.GetString("s","")
	this.Data["s"] = search
	filters := make(map[string] interface{})
	if !(search == ""){
		filters["text__icontains"]=search
	}
	roles,total := m.MenuPageList(page,size,filters)

	paginator := utils.NewPaginator(this.Ctx.Request, size, total)
	this.Data["data"] = roles
	this.Data["paginator"] = paginator
	this.Layout = layout
	this.TplName = "menu/list.html"
}

func (this *MenuController) Add()  {
	if this.IsGet(){
		this.Layout = layout
		this.TplName = "menu/add.html"
	}
}
