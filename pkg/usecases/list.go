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

func (usecases *NotesListUsecases) Create(userId int, list model.NotesList) (int, error) {
	return usecases.repository.Create(userId, list)
}
func (usecases *NotesListUsecases) GetAllList(userId int) ([]model.NotesList, error) {
	return usecases.repository.GetAll(userId)
}
func (usecases *NotesListUsecases) GetListById(userId, listId int) (model.NotesList, error) {
	return usecases.repository.GetListById(userId, listId)
}
func (usecases *NotesListUsecases) Update(userId, listId int, list model.UpdateListInput) error {
	if err := list.Validate(); err != nil {
		return err
	}
	return usecases.repository.Update(userId, listId, list)
}
func (usecases *NotesListUsecases) Delete(userId, listId int) error {
	return usecases.repository.Delete(userId, listId)
}
