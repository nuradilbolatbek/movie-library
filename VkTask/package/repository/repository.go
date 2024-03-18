package repository

import (
	"VkTask"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user VkTask.User) (int, error)
	GetUser(username, password string) (VkTask.User, error)
}

type ActorManagement interface {
	CreateActor(actor VkTask.Actor) (int, error)
	UpdateActor(id int, actor VkTask.Actor) error
	DeleteActor(id int) error
	GetActorByID(id int) (VkTask.Actor, error)
	GetAllActors() ([]VkTask.Actor, error)
	SearchActors(nameFragment string) ([]VkTask.Actor, error)
}

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

// Repository struct aggregates all repositories
type Repository struct {
	Authorization
	ActorManagement
	MovieManagement
}

// NewRepository creates a new Repository instance
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:   NewAuthPostgresRepo(db),
		ActorManagement: NewActorPostgresRepo(db),
		MovieManagement: NewMoviePostgresRepo(db),
	}
}
