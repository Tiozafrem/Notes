package model

import "errors"

// Notes model
type NotesList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

// Update notes in model
type ListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// Check update notes for not null all value
func (item ListInput) Validate() error {
	if item.Title == nil && item.Description == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
