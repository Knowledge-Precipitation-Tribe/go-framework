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
