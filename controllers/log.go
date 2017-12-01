package controllers

import (
	m "flat-goadmin/models"
)

type LogController struct {
	BaseController
}


func (this *LogController) Index()  {
	page,_ := this.GetInt("p",1)
	search := this.GetString("s","")

	this.Data["s"] = search
	filters := make(map[string] interface{})
	if !(search == ""){
		filters["log_title__icontains"]=search
	}
	data,total := m.LogPageList(page,size,filters)
	this.renderPageTpl(data,total,"log/list.html")
}