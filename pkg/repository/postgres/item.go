package postgres

import (
	"fmt"
	"notes/model"
	"strings"

	"github.com/jmoiron/sqlx"
)

type NotesItemPostgres struct {
	db *sqlx.DB
}

func NewNotesItemPostgres(db *sqlx.DB) *NotesItemPostgres {
	return &NotesItemPostgres{db: db}
}

func (r *NotesItemPostgres) Create(userId, listId int, item model.NoteItem) (int, error) {
	transanction, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf(
		`INSERT INTO %s (title, description, done) values ($1, $2, $3) RETURNING id`,
		noteItemsTable)

	row := transanction.QueryRow(createItemQuery, item.Title, item.Description, item.Done)
	if err := row.Scan(&itemId); err != nil {
		transanction.Rollback()
		return 0, err
	}

	createListQuery := fmt.Sprintf(
		`INSERT INTO %s (list_id, item_id) values ($1, $2)`,
		listsItemsTable)
	_, err = transanction.Exec(createListQuery, listId, itemId)
	if err != nil {
		transanction.Rollback()
		return 0, err
	}

	return itemId, transanction.Commit()
}

func (r *NotesItemPostgres) GetAll(userId, listId int) ([]model.NoteItem, error) {
	var items []model.NoteItem
	query := fmt.Sprintf(
		`SELECT item_table.id, item_table.title, item_table.description, item_table.done
		FROM %s item_table
		INNER JOIN %s list_item_table ON
			item_table.id = list_item_table.item_id
		INNER JOIN %s list_user_table ON
			list_item_table.list_id = list_user_table.list_id
		WHERE list_user_table.user_id = $1 AND
			list_item_table.list_id = $2`,
		noteItemsTable, listsItemsTable, usersListTable)
	if err := r.db.Select(&items, query, userId, listId); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *NotesItemPostgres) GetItemById(userId, itemId int) (model.NoteItem, error) {
	var item model.NoteItem
	query := fmt.Sprintf(
		`SELECT item_table.id, item_table.title, item_table.description, item_table.done 
		FROM %s item_table
		INNER JOIN %s list_item_table ON
			item_table.id = list_item_table.item_id
		INNER JOIN %s list_user_table ON
			list_item_table.list_id = list_user_table.list_id
		WHERE list_user_table.user_id = $1 AND
			item_table.id = $2 `,
		noteItemsTable, listsItemsTable, usersListTable)
	if err := r.db.Get(&item, query, userId, itemId); err != nil {
		return item, err
	}
	return item, nil
}

func (r *NotesItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s item_table
		USING %s list_item_table, %s list_user_table
		WHERE item_table.id = list_item_table.item_id AND
			list_item_table.list_id = list_user_table.list_id AND
			list_user_table.user_id = $1 AND
			item_table.id = $2`,
		noteItemsTable, listsItemsTable, usersListTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *NotesItemPostgres) Update(userId, itemId int, item model.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if item.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *item.Title)
		argId++
	}

	if item.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *item.Description)
		argId++
	}

	if item.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *item.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		`UPDATE %s item_table
		SET %s
		FROM %s list_item_table, %s list_user_table
		WHERE item_table.id = list_item_table.item_id AND
		list_item_table.list_id = list_user_table.list_id AND
		list_user_table.user_id = $%d AND
		item_table.id = $%d`,
		noteItemsTable, setQuery, listsItemsTable, usersListTable,
		argId, argId+1)
	args = append(args, userId, itemId)
	_, err := r.db.Exec(query, args...)
	return err
}
