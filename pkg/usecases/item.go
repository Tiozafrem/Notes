package usecases

import (
	"fmt"
	"notes/model"
	"notes/pkg/hub"
	"notes/pkg/repository"
)

type NotesItemUsecases struct {
	hub            hub.HubNotify
	repository     repository.NoteItem
	repositoryList repository.NoteList
}

func NewNotesItemUsecases(repository repository.NoteItem, repositoryList repository.NoteList,
	hub hub.HubNotify) *NotesItemUsecases {
	return &NotesItemUsecases{repository: repository, repositoryList: repositoryList, hub: hub}
}

func (u *NotesItemUsecases) Create(userId, listId int, item model.NoteItem) (int, error) {
	_, err := u.repositoryList.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}

	id, err := u.repository.Create(userId, listId, item)
	if err != nil {
		return 0, err
	}

	go func() {
		users, _ := u.repositoryList.GetAllUserListByListId(listId)
		for _, user := range users {
			u.hub.EmitUser(user.UserId, fmt.Sprintf("update list %d", listId))
		}
	}()
	return id, nil
}

func (u *NotesItemUsecases) GetAll(userId, listId int) ([]model.NoteItem, error) {
	return u.repository.GetAll(userId, listId)
}

func (u *NotesItemUsecases) GetItemById(userId, itemId int) (model.NoteItem, error) {
	return u.repository.GetItemById(userId, itemId)
}

func (u *NotesItemUsecases) Delete(userId, itemId int) error {
	err := u.repository.Delete(userId, itemId)
	if err != nil {
		return err
	}

	go func() {
		users, _ := u.repository.GetAllUserListByItemId(itemId)
		for _, user := range users {
			u.hub.EmitUser(user.UserId, fmt.Sprintf("delete item %d", itemId))
		}
	}()

	return nil
}

func (u *NotesItemUsecases) Update(userId, itemId int, item model.ItemInput) error {
	if err := item.Validate(); err != nil {
		return err
	}

	err := u.repository.Update(userId, itemId, item)
	if err != nil {
		return err
	}

	go func() {
		users, _ := u.repository.GetAllUserListByItemId(itemId)
		for _, user := range users {
			u.hub.EmitUser(user.UserId, fmt.Sprintf("update item %d", itemId))
		}
	}()

	return nil
}
