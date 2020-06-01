package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris-demo/web/controllers"
)

func main() {
	app := iris.New()
	//设置日志级别
	app.Logger().SetLevel("debug")
	//注册视图模版, HTML函数的两个参数分别是视图所在位置，以及视图文件的拓展名
	app.RegisterView(iris.HTML("./web/views", ".html"))
	//注册控制器
	//Party下存放URL映射，Handle配置对应的控制器
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))

	app.Run(
		iris.Addr("localhost:8080"),
	)
}
