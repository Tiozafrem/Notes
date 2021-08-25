package model

import "errors"

type NoteItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (item UpdateItemInput) Validate() error {
	if item.Title == nil && item.Description == nil && item.Done == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
