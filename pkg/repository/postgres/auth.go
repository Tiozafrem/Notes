package postgres

import (
	"fmt"
	"notes/model"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash, password_salt) values ($1, $2, $3, $4) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password, user.Salt)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id, password_hash, password_salt FROM %s WHERE username=$1", userTable)
	err := r.db.Get(&user, query, username)

	return user, err
}
