# Iris MVC案例

## 项目结构

![](../../.gitbook/assets/image%20%287%29.png)

## datamodels

```go
package datamodels

type Movie struct {
	Name string
}
```

## repositories

```go
package repositories

import "iris-demo/datamodels"

type MovieRepository interface {
	GetMovieName() string
}

type MovieRepositoryManager struct {
}

func NewMovieRepositoryManager() MovieRepository {
	return &MovieRepositoryManager{}
}

func (m MovieRepositoryManager) GetMovieName() string {
	//todo:根据数据库返回
	movie := &datamodels.Movie{
		Name: "test",
	}
	return movie.Name
}

```

## services

```go
package services

import (
	"fmt"
	"iris-demo/repositories"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct{
	repo repositories.MovieRepository
}

func NewMovieServiceManager(repo repositories.MovieRepository) MovieService{
	return &MovieServiceManager{
		repo:repo,
	}
}

func (m *MovieServiceManager) ShowMovieName() string{
	fmt.Println("movie name : ", m.repo.GetMovieName())
	return "movie name : " + m.repo.GetMovieName()
}
```

## controllers

```go
package controllers

import (
	"github.com/kataras/iris/mvc"
	"iris-demo/repositories"
	"iris-demo/services"
)

type MovieController struct {
}

//get请求
func (c *MovieController) Get() mvc.View {
	movieRepository := repositories.NewMovieRepositoryManager()
	movie := services.NewMovieServiceManager(movieRepository)
	movieName := movie.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: movieName,
	}
}

```

## views/movie

```markup
<h2>{{.}}</h2>
```

## main.go

```go
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
```

## 代码位置

Github：[iris-demo](https://github.com/Knowledge-Precipitation-Tribe/go-framework/tree/master/code/iris-demo)

