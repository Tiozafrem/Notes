package usecases

import (
	"fmt"
	"notes/model"
	"notes/pkg/hub"
	"notes/pkg/repository"
)

type NotesListUsecases struct {
	hub        hub.HubNotify
	repository repository.NoteList
}

func NewNotesListUsecases(repository repository.NoteList, hub hub.HubNotify) *NotesListUsecases {
	return &NotesListUsecases{repository: repository, hub: hub}
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
func (u *NotesListUsecases) Update(userId, listId int, list model.ListInput) error {
	if err := list.Validate(); err != nil {
		return err
	}

	go func() {
		users, _ := u.repository.GetAllUserListByListId(listId)
		for _, user := range users {
			u.hub.EmitUser(user.UserId, fmt.Sprintf("update list %d", listId))
		}
	}()
	return u.repository.Update(userId, listId, list)
}
func (u *NotesListUsecases) Delete(userId, listId int) error {
	err := u.repository.Delete(userId, listId)
	if err != nil {
		return err
	}

	go func() {
		users, _ := u.repository.GetAllUserListByListId(listId)
		for _, user := range users {
			u.hub.EmitUser(user.UserId, fmt.Sprintf("delete list %d", listId))
		}
	}()

	return nil
}
