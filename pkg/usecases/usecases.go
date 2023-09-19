package usecases

import (
	"notes/model"
	"notes/pkg/repository"
	"notes/pkg/usecases/auth"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password, nameDevice string) (auth.Tokens, error)
	ParseTokenToUserId(toen string) (int, error)
	RefreshToken(refresh_token string) (auth.Tokens, error)
	NewRefreshToken() (string, error)
	NewAccessToken(deviceId int) (string, error)
}

type NoteList interface {
	Create(userId int, list model.NotesList) (int, error)
	GetAllList(userId int) ([]model.NotesList, error)
	GetListById(userId, listId int) (model.NotesList, error)
	Update(userId, listId int, list model.ListInput) error
	Delete(userId, listId int) error
}

type NoteItem interface {
	Create(userId, listId int, item model.NoteItem) (int, error)
	GetAll(userId, listId int) ([]model.NoteItem, error)
	GetItemById(userId, itemId int) (model.NoteItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, item model.ItemInput) error
}

type Usecases struct {
	Authorization
	NoteList
	NoteItem
}

func NewUsecases(repositorys *repository.Repository) *Usecases {
	return &Usecases{
		Authorization: auth.NewAuthUsecases(repositorys.Authorization),
		NoteList:      NewNotesListUsecases(repositorys.NoteList),
		NoteItem:      NewNotesItemUsecases(repositorys.NoteItem, repositorys.NoteList),
	}
}
