package controllers

import (
	m "flat-goadmin/models"
	"github.com/beego/wetalk/modules/utils"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	BaseController
}

func (this *UserController) Index()  {
	page,err:= this.GetInt("p")
	if err != nil{
		page = 1;
	}
	s := this.GetString("s")
	//查询条件
	filters := make(map[string] interface{})
	if s != ""{
		filters["user_name__icontains"] = s
		this.Data["s"] = s
	}
	users,total := m.PageList(page,size,filters)
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
	id,_:= this.GetInt64("id")
	_,err := orm.NewOrm().Delete(&m.User{Id:id})
	if err == nil{
		this.Redirect("/user",302)
	}
}

func (this *UserController) Edit()  {
	id,_ := this.GetInt64("id")
	user := m.User{Id:id}
	if this.IsGet(){
		orm.NewOrm().Read(&user);
		this.Data["user"] = user
		this.Layout = layout
		this.TplName = "user/edit.html"
	}else {
		this.ParseForm(&user)
		if _,err := m.UpdateUser(&user);err == nil{
			this.Redirect("/user",302)
		}else {
			this.Rsp(false,err.Error())
		}
	}
}