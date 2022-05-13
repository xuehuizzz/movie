package controllers

import "github.com/astaxie/beego"
/**
  该控制器处理页面错误请求
*/
type ErrorController struct {
	beego.Controller
}


func (e *ErrorController) Error404() {
	e.Data["content"] = "很抱歉您访问的地址或者方法不存在"
	e.TplName = "404.html"
}
