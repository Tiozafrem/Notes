package usecases

import (
	"notes/model"
	"notes/pkg/repository"
	"notes/pkg/usecases/auth"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseTokenToUserId(toen string) (int, error)
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

type Usecases struct {
	Authorization
	NoteList
	NoteItem
}

func NewUsecases(repository *repository.Repository) *Usecases {
	return &Usecases{
		Authorization: auth.NewAuthUsecases(repository.Authorization),
		NoteList:      NewNotesListUsecases(repository.NoteList),
	}
}
