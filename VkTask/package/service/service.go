package service

import (
	"VkTask"
	"VkTask/package/repository"
)

type Authorization interface {
	CreateUser(user VkTask.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, string, error)
}

// ActorManagement interface for managing actors
type ActorManagement interface {
	CreateActor(actor VkTask.Actor) (int, error)
	UpdateActor(id int, actor VkTask.Actor) error
	DeleteActor(id int) error
	GetActorByID(id int) (VkTask.Actor, error)
	GetAllActors() ([]VkTask.Actor, error)
	SearchActors(nameFragment string) ([]VkTask.Actor, error)
}

// MovieManagement interface for managing movies
type MovieManagement interface {
	CreateMovie(movie VkTask.Movie) (int, error)
	UpdateMovie(id int, movie VkTask.Movie) error
	DeleteMovie(id int) error
	GetMovieByID(id int) (VkTask.Movie, error)
	GetAllMovies() ([]VkTask.Movie, error)
	AddActorToMovie(movieID, actorID int) error
	RemoveActorFromMovie(movieID, actorID int) error
	SearchMovies(titleFragment, actorNameFragment string) ([]VkTask.Movie, error)
	GetAllMoviesSorted(sortParam string) ([]VkTask.Movie, error)
}

// Service struct to aggregate all services
type Service struct {
	Authorization
	ActorManagement
	MovieManagement
}

// NewService creates a new Service instance
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:   NewAuthService(repos.Authorization),
		ActorManagement: NewActorService(repos.ActorManagement),
		MovieManagement: NewMovieService(repos.MovieManagement),
	}
}
