package service

import (
	"VkTask"
	"VkTask/package/repository"
)

type MovieManagementService struct {
	repo repository.MovieManagement
}

func NewMovieService(repo repository.MovieManagement) *MovieManagementService {
	return &MovieManagementService{
		repo: repo,
	}
}

// CreateMovie creates a new movie
func (s *MovieManagementService) CreateMovie(movie VkTask.Movie) (int, error) {
	return s.repo.CreateMovie(movie)
}

// UpdateMovie updates an existing movie by its ID
func (s *MovieManagementService) UpdateMovie(id int, movie VkTask.Movie) error {
	return s.repo.UpdateMovie(id, movie)
}

// DeleteMovie deletes a movie by its ID
func (s *MovieManagementService) DeleteMovie(id int) error {
	return s.repo.DeleteMovie(id)
}

// GetMovieByID retrieves a movie by its ID
func (s *MovieManagementService) GetMovieByID(id int) (VkTask.Movie, error) {
	return s.repo.GetMovieByID(id)
}

// GetAllMovies retrieves all movies
func (s *MovieManagementService) GetAllMovies() ([]VkTask.Movie, error) {
	return s.repo.GetAllMovies()
}

// AddActorToMovie adds an association between a movie and an actor
func (s *MovieManagementService) AddActorToMovie(movieID, actorID int) error {
	return s.repo.AddActorToMovie(movieID, actorID)
}

// RemoveActorFromMovie removes an association between a movie and an actor
func (s *MovieManagementService) RemoveActorFromMovie(movieID, actorID int) error {
	return s.repo.RemoveActorFromMovie(movieID, actorID)
}

func (s *MovieManagementService) SearchMovies(titleFragment, actorNameFragment string) ([]VkTask.Movie, error) {
	return s.repo.SearchMovies(titleFragment, actorNameFragment)
}

func (s *MovieManagementService) GetAllMoviesSorted(sortParam string) ([]VkTask.Movie, error) {
	return s.repo.GetAllMoviesSorted(sortParam)
}
