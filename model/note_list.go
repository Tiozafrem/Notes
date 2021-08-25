package model

import "errors"

type NotesList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (item UpdateListInput) Validate() error {
	if item.Title == nil && item.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
