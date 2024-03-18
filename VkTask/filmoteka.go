package VkTask

import "time"

type Actor struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" binding:"required"`
	Gender    string    `json:"gender" binding:"required"`
	Birthdate time.Time `json:"birthdate" db:"birthdate"`
}

type Movie struct {
	Id          int       `json:"id" db:"id"`
	Title       string    `json:"title" binding:"required,max=150"`
	Description string    `json:"description" binding:"max=1000"`
	ReleaseDate time.Time `json:"release_date" db:"release_date"`
	Rating      float32   `json:"rating" binding:"gte=0,lte=10"`
	ActorsID    int       `json:"actors" db:"-"`
}

type MovieActor struct {
	MovieID int `db:"movie_id"`
	ActorID int `db:"actor_id"`
}
