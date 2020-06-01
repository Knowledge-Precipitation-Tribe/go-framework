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
