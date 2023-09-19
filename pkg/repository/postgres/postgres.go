package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	userTable       = "users"
	noteListTable   = "lists"
	usersListTable  = "users_lists"
	noteItemsTable  = "items"
	listsItemsTable = "lists_items"
	deviceUserTable = "devices_users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	return db, nil
}
