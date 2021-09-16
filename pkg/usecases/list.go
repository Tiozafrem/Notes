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

func (u *NotesListUsecases) Create(userId int, list model.NotesList) (int, error) {
	return u.repository.Create(userId, list)
}
func (u *NotesListUsecases) GetAllList(userId int) ([]model.NotesList, error) {
	return u.repository.GetAll(userId)
}
func (u *NotesListUsecases) GetListById(userId, listId int) (model.NotesList, error) {
	return u.repository.GetListById(userId, listId)
}
func (u *NotesListUsecases) Update(userId, listId int, list model.UpdateListInput) error {
	if err := list.Validate(); err != nil {
		return err
	}
	return u.repository.Update(userId, listId, list)
}
func (u *NotesListUsecases) Delete(userId, listId int) error {
	return u.repository.Delete(userId, listId)
}
