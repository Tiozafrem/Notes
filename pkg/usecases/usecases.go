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
	}
}
