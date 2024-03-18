package repository

import (
	"VkTask"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type MoviePostgres struct {
	db *sqlx.DB
}

func NewMoviePostgresRepo(db *sqlx.DB) *MoviePostgres {
	return &MoviePostgres{db: db}
}

func (r *MoviePostgres) CreateMovie(movie VkTask.Movie) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, description, release_date, rating) VALUES ($1, $2, $3, $4) RETURNING id", moviesTable)
	err := r.db.QueryRow(query, movie.Title, movie.Description, movie.ReleaseDate, movie.Rating).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MoviePostgres) UpdateMovie(id int, movie VkTask.Movie) error {
	query := fmt.Sprintf("UPDATE %s SET title = $1, description = $2, release_date = $3, rating = $4 WHERE id = $5", moviesTable)
	_, err := r.db.Exec(query, movie.Title, movie.Description, movie.ReleaseDate, movie.Rating, id)
	return err
}

func (r *MoviePostgres) DeleteMovie(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", moviesTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *MoviePostgres) GetMovieByID(id int) (VkTask.Movie, error) {
	var movie VkTask.Movie
	query := fmt.Sprintf("SELECT id, title, description, release_date, rating FROM %s WHERE id = $1", moviesTable)
	err := r.db.Get(&movie, query, id)
	return movie, err
}

func (r *MoviePostgres) GetAllMovies() ([]VkTask.Movie, error) {
	var movies []VkTask.Movie
	query := fmt.Sprintf("SELECT id, title, description, release_date, rating FROM %s", moviesTable)
	err := r.db.Select(&movies, query)
	return movies, err
}

func (r *MoviePostgres) AddActorToMovie(movieID, actorID int) error {
	query := fmt.Sprintf("INSERT INTO %s (movie_id, actor_id) VALUES ($1, $2)", moviesActorsTable)
	_, err := r.db.Exec(query, movieID, actorID)
	return err
}

func (r *MoviePostgres) RemoveActorFromMovie(movieID, actorID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE movie_id = $1 AND actor_id = $2", moviesActorsTable)
	_, err := r.db.Exec(query, movieID, actorID)
	return err
}
func (r *MoviePostgres) SearchMovies(titleFragment, actorNameFragment string) ([]VkTask.Movie, error) {
	var movies []VkTask.Movie
	var err error // Declare err here

	// Assuming you have a JOIN query if searching by actor name
	query := `SELECT m.id, m.title, m.description, m.release_date, m.rating FROM movies AS m`
	if actorNameFragment != "" {
		query += ` JOIN movies_actors AS ma ON m.id = ma.movie_id`
		query += ` JOIN actors AS a ON ma.actor_id = a.id`
		query += ` WHERE a.name ILIKE '%' || $1 || '%' OR m.title ILIKE '%' || $2 || '%'`
		err = r.db.Select(&movies, query, actorNameFragment, titleFragment)
	} else {
		query += ` WHERE m.title ILIKE '%' || $1 || '%'`
		err = r.db.Select(&movies, query, titleFragment)
	}
	return movies, err
}

func (r *MoviePostgres) GetAllMoviesSorted(sortParam string) ([]VkTask.Movie, error) {
	var movies []VkTask.Movie
	var query string

	switch sortParam {
	case "title":
		query = "SELECT * FROM movies ORDER BY title"
	case "release_date":
		query = "SELECT * FROM movies ORDER BY release_date"
	default: // default to sort by rating in descending order
		query = "SELECT * FROM movies ORDER BY rating DESC"
	}

	err := r.db.Select(&movies, query)
	return movies, err
}
