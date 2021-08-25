package usecases

import "notes/pkg/repository"

type Authorization interface {
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
	return &Usecases{}
}
