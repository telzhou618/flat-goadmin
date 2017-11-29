package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/wetalk/modules/utils"
)

var size int = 15 //默认页大小，取app.conf文件中的page_size

const (
	layout string = "common/layout.html"
)

type BaseController struct {
	beego.Controller
}

//初始化基础数据，启动执行
func init() {
	 size,_= beego.AppConfig.Int("page_size")
}

//设置默认模板，每次请求执行
func (c *BaseController) Prepare()  {
	//全局验证处理

}

//转JSON
func (c *BaseController)Rsp(success bool,str interface{})  {
	c.Data["json"] = &map[string]interface{}{"success":success,"data" : str}
	c.ServeJSON()
	c.StopRun()
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

//重定向302
func (c *BaseController) Redirect302(url string) {
	c.Redirect(url,302)
	return
}

//渲染到布局+模板
func (c *BaseController) renderLayTpl(layout,url string) {
	c.Layout = layout
	c.TplName = url
	return
}
//渲染到模板
func (c *BaseController) renderTpl(url string) {
	c.Layout = layout
	c.TplName = url
	return
}

//渲染分页列表
/**
 size 页大小
  total 总记录数
 */
func (c *BaseController) renderPageTpl(data interface{},total int64,url string) {
	paginator := utils.NewPaginator(c.Ctx.Request, size, total)
	c.Data["data"] = data
	c.Data["paginator"] = paginator
	c.Layout = layout
	c.TplName = url
	return
}