package controllers

import (
	"github.com/astaxie/beego"
)

const (
	layout string = "common/layout.html"
	size int = 10
)

type BaseController struct {
	beego.Controller
}

//转JSON
func (c *BaseController)Rsp(success bool,str string)  {
	c.Data["json"] = &map[string]interface{}{"success":success,"data" : str}
	c.ServeJSON()
}

//是否为GET请求
func (c *BaseController) IsGet() bool {

	if mt := c.Ctx.Request.Method; mt == "GET" {
		return true
	}else {
		return false
	}
}

// 是否为POST请求
func (c *BaseController) IsPOST() bool {

	if mt := c.Ctx.Request.Method; mt == "POST" {
		return true
	}else {
		return false
	}
}