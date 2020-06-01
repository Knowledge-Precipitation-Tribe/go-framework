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