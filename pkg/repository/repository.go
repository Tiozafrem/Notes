package repository

import (
	"notes/model"
	"notes/pkg/repository/postgres"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type NoteList interface {
	CreateList(userId int, list model.NotesList) (int, error)
	GetAllListUserId(userId int) ([]model.NotesList, error)
	GetListByIdUserId(userId, listId int) (model.NotesList, error)
	UpdateListByIdUserId(userId, listId int, list model.UpdateListInput) error
	DeleteListByIdUserId(userId, listId int) error
}

type NoteItem interface {
}

type Repository struct {
	Authorization
	NoteList
	NoteItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		NoteList:      postgres.NewNotesListPostgres(db),
		NoteItem:      nil,
	}
}
