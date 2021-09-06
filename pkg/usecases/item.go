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

func (usecases *NotesItemUsecases) Create(userId, listId int, item model.NoteItem) (int, error) {
	_, err := usecases.repositoryList.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}
	return usecases.repository.Create(userId, listId, item)
}

func (usecases *NotesItemUsecases) GetAll(userId, listId int) ([]model.NoteItem, error) {
	return usecases.repository.GetAll(userId, listId)
}

func (usecases *NotesItemUsecases) GetItemById(userId, itemId int) (model.NoteItem, error) {
	return usecases.repository.GetItemById(userId, itemId)
}

func (usecases *NotesItemUsecases) Delete(userId, itemId int) error {
	return usecases.repository.Delete(userId, itemId)
}
func (usecases *NotesItemUsecases) Update(userId, itemId int, item model.UpdateItemInput) error {
	if err := item.Validate(); err != nil {
		return err
	}
	return usecases.repository.Update(userId, itemId, item)
}
