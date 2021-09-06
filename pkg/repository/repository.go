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
	Create(userId int, list model.NotesList) (int, error)
	GetAll(userId int) ([]model.NotesList, error)
	GetListById(userId, listId int) (model.NotesList, error)
	Update(userId, listId int, list model.UpdateListInput) error
	Delete(userId, listId int) error
}

type NoteItem interface {
	Create(userId, listId int, item model.NoteItem) (int, error)
	GetAll(userId, listId int) ([]model.NoteItem, error)
	GetItemById(userId, itemId int) (model.NoteItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, item model.UpdateItemInput) error
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
		NoteItem:      postgres.NewNotesItemPostgres(db),
	}
}
