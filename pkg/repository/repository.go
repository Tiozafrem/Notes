package repository

type Authorization interface {
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

func NewRepository() *Repository {
	return &Repository{}
}
