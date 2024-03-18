package repository

import (
	"VkTask"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ActorPostgres struct {
	db *sqlx.DB
}

func NewActorPostgresRepo(db *sqlx.DB) *ActorPostgres {
	return &ActorPostgres{db: db}
}

// CreateActor adds a new actor to the database
func (r *ActorPostgres) CreateActor(actor VkTask.Actor) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, gender, birth_date) VALUES ($1, $2, $3) RETURNING id", actorsTable)
	err := r.db.QueryRow(query, actor.Name, actor.Gender, actor.Birthdate).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// UpdateActor updates an existing actor's information
func (r *ActorPostgres) UpdateActor(id int, actor VkTask.Actor) error {
	query := fmt.Sprintf("UPDATE %s SET name = $1, gender = $2, birth_date = $3 WHERE id = $4", actorsTable)
	_, err := r.db.Exec(query, actor.Name, actor.Gender, actor.Birthdate, id)
	return err
}

// DeleteActor removes an actor from the database
func (r *ActorPostgres) DeleteActor(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", actorsTable)
	_, err := r.db.Exec(query, id)
	return err
}

// GetActorByID retrieves an actor by their ID
func (r *ActorPostgres) GetActorByID(id int) (VkTask.Actor, error) {
	var actor VkTask.Actor
	query := fmt.Sprintf("SELECT id, name, gender, birth_date FROM %s WHERE id = $1", actorsTable)
	err := r.db.Get(&actor, query, id)
	return actor, err
}

// GetAllActors retrieves all actors from the database
func (r *ActorPostgres) GetAllActors() ([]VkTask.Actor, error) {
	var actors []VkTask.Actor
	query := fmt.Sprintf("SELECT id, name, gender, birth_date FROM %s", actorsTable)
	err := r.db.Select(&actors, query)
	return actors, err
}

func (r *ActorPostgres) SearchActors(nameFragment string) ([]VkTask.Actor, error) {
	var actors []VkTask.Actor
	query := `SELECT id, name, gender, birthdate FROM actors WHERE name ILIKE '%' || $1 || '%'`
	err := r.db.Select(&actors, query, nameFragment)
	return actors, err
}
