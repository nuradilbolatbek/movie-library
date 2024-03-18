package repository

import (
	"VkTask"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthorizationPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgresRepo(db *sqlx.DB) *AuthorizationPostgres {
	return &AuthorizationPostgres{db: db}
}

func (r *AuthorizationPostgres) CreateUser(user VkTask.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash, role) values ($1, $2, $3, $4) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthorizationPostgres) GetUser(username, passwordHash string) (VkTask.User, error) {
	var user VkTask.User
	query := fmt.Sprintf("SELECT id, name, username, role FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, passwordHash)

	return user, err
}
