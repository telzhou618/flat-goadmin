package controllers

import (
	m "flat-goadmin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MenuController struct {
	BaseController
}


func (this *MenuController) Index()  {
	beego.ReadFromRequest(&this.Controller)
	page,_ := this.GetInt("p",1)
	search := this.GetString("s","")

	this.Data["s"] = search
	filters := make(map[string] interface{})
	if !(search == ""){
		filters["text__icontains"]=search
	}
	data,total := m.MenuPageList(page,size,filters)
	this.renderPageTpl(data,total,"menu/list.html")
}

func (this *MenuController) Add()  {
	if this.IsGet(){
		this.Layout = layout
		this.TplName = "menu/add.html"
	}else {
		menu := m.Menu{}
		this.ParseForm(&menu)
		orm.NewOrm().Insert(&menu)
		flash := beego.NewFlash()
		flash.Success("保存成功！")
		flash.Store(&this.Controller)
		this.Redirect302("/menu")
	}
}
