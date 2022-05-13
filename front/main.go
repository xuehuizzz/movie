package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"movie/front/controllers"
	_ "movie/front/routers"
	"movie/front/utils"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	orm.RunCommand()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.InsertFilter("/main/*",beego.BeforeRouter,utils.LoginFilter)   // 未登录请求拦截
	logs.SetLogger(logs.AdapterMultiFile,`{"filename":"logs/movie.log","separate":["error","info"]}`)  // 日志
	logs.SetLogFuncCallDepth(3)
	beego.SetStaticPath("/upload","upload")
	beego.Run()
}
