# Iris MVC

## 工作流程

![](../../.gitbook/assets/image%20%288%29.png)

## 目录结构

![](../../.gitbook/assets/image%20%287%29.png)

* datamodels:用于存放一些模型，主要是存放一些结构体
* repositories:用于与持久化存储进行交互
* services:用于处理业务逻辑
* controllers:用于存放控制器，也就是你的请求调用哪些处理逻辑进行处理
* views:用于存放视图，像html，js的一些东西
* main.go:程序的主入口

接下来一下main.go文件

```go
package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	//设置日志级别
	app.Logger().SetLevel("debug")
	//注册视图模版, HTML函数的两个参数分别是视图所在位置，以及视图文件的拓展名
	app.RegisterView(iris.HTML("./web/views", ".html"))
	//注册控制器
	
	app.Run(
		iris.Addr("localhost:8080"), 
	)
}
```

