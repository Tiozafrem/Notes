package model

import "errors"

// Item in notes model
type NoteItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

// Update item in notes model
type ItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

// Check update item for not null all value
func (item ItemInput) Validate() error {
	if item.Title == nil && item.Description == nil && item.Done == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
