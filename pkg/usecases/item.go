package usecases

import (
	"notes/model"
	"notes/pkg/repository"
)

type NotesItemUsecases struct {
	repository     repository.NoteItem
	repositoryList repository.NoteList
}

func NewNotesItemUsecases(repository repository.NoteItem, repositoryList repository.NoteList) *NotesItemUsecases {
	return &NotesItemUsecases{repository: repository, repositoryList: repositoryList}
}

func (u *NotesItemUsecases) Create(userId, listId int, item model.NoteItem) (int, error) {
	_, err := u.repositoryList.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}
	return u.repository.Create(userId, listId, item)
}

func (u *NotesItemUsecases) GetAll(userId, listId int) ([]model.NoteItem, error) {
	return u.repository.GetAll(userId, listId)
}

func (u *NotesItemUsecases) GetItemById(userId, itemId int) (model.NoteItem, error) {
	return u.repository.GetItemById(userId, itemId)
}

func (u *NotesItemUsecases) Delete(userId, itemId int) error {
	return u.repository.Delete(userId, itemId)
}
func (u *NotesItemUsecases) Update(userId, itemId int, item model.ItemInput) error {
	if err := item.Validate(); err != nil {
		return err
	}
	return u.repository.Update(userId, itemId, item)
}
