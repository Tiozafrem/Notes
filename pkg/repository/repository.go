package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
	CreateUser()
}

type NoteList interface {
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
		Authorization: nil,
		NoteList:      nil,
		NoteItem:      nil,
	}
}
