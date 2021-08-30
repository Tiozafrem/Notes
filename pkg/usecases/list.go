package usecases

import (
	"notes/model"
	"notes/pkg/repository"
)

type NotesListUsecases struct {
	repository repository.NoteList
}

func NewNotesListUsecases(repository repository.NoteList) *NotesListUsecases {
	return &NotesListUsecases{repository: repository}
}

func (usecases *NotesListUsecases) CreateList(userId int, list model.NotesList) (int, error) {
	return usecases.repository.CreateList(userId, list)
}
func (usecases *NotesListUsecases) GetAllListUserId(userId int) ([]model.NotesList, error) {
	return usecases.repository.GetAllListUserId(userId)
}
func (usecases *NotesListUsecases) GetListByIdUserId(userId, listId int) (model.NotesList, error) {
	return usecases.repository.GetListByIdUserId(userId, listId)
}
func (usecases *NotesListUsecases) UpdateListByIdUserId(userId, listId int, list model.UpdateListInput) error {
	if err := list.Validate(); err != nil {
		return err
	}
	return usecases.repository.UpdateListByIdUserId(userId, listId, list)
}
func (usecases *NotesListUsecases) DeleteListByIdUserId(userId, listId int) error {
	return usecases.repository.DeleteListByIdUserId(userId, listId)
}
