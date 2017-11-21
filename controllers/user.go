package controllers

import (
	m "flat-goadmin/models"
	"github.com/beego/wetalk/modules/utils"
)

type UserController struct {
	BaseController
}

func (this *UserController) Index()  {
	page,err:= this.GetInt("p")
	if err != nil{
		page = 1;
	}
	users,total := m.PageList(page,size)
	paginator := utils.NewPaginator(this.Ctx.Request, size, total)
	this.Data["data"] = users
	this.Data["paginator"] = paginator
	this.Layout = layout
	this.TplName = "user/list.html"
}
func (this *UserController) Add()  {
	this.Layout = layout
	if this.IsGet(){
		this.TplName = "user/add.html"
		return
	}else {
		user := m.User{}
		if err := this.ParseForm(&user); err != nil{
			this.TplName = "user/add.html"
			this.Data["error"] = err.Error()
			return
		}else {
			if _,err := m.AddUser(&user); err !=nil{
				this.TplName = "user/add.html"
				this.Data["error"] = err.Error()
				return
			}else {
				this.Redirect("/user",302)
				return
			}
		}
	}
}

func (this *UserController) Del()  {
	this.Rsp(true,"Del")
}

func (this *UserController) Update()  {
	this.Rsp(true,"Update")
}

func (this *UserController) Find()  {
	this.Rsp(true,"Find")
}