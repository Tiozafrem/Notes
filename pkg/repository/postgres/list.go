package postgres

import (
	"fmt"
	"notes/model"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type NotesListPostgres struct {
	db *sqlx.DB
}

func NewNotesListPostgres(db *sqlx.DB) *NotesListPostgres {
	return &NotesListPostgres{db: db}
}

func (r *NotesListPostgres) Create(userId int, list model.NotesList) (int, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var listId int
	createListQuery := fmt.Sprintf(
		`INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id`,
		noteListTable,
	)
	row := transaction.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&listId); err != nil {
		transaction.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf(
		`INSERT INTO %s (user_id, list_id) VALUES ($1, $2)`,
		usersListTable,
	)
	_, err = transaction.Exec(createUsersListQuery, userId, listId)
	if err != nil {
		transaction.Rollback()
		return 0, err
	}
	return listId, transaction.Commit()
}
func (r *NotesListPostgres) GetAll(userId int) ([]model.NotesList, error) {
	var lists []model.NotesList

	query := fmt.Sprintf(
		`SELECT list_table.id, list_table.title, list_table.description 
		FROM %s list_table 
		INNER JOIN %s list_user_table on 
			list_table.id = list_user_table.list_id 
		WHERE list_user_table.user_id = $1`,
		noteListTable, usersListTable,
	)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}
func (r *NotesListPostgres) GetListById(userId, listId int) (model.NotesList, error) {
	var list model.NotesList

	query := fmt.Sprintf(
		`SELECT list_table.id, list_table.title, list_table.description
		FROM %s list_table
		INNER JOIN %s list_user_table on
			list_table.id = list_user_table.list_id
		WHERE list_user_table.user_id = $1 AND
			list_table.id = $2`,
		noteListTable, usersListTable,
	)
	err := r.db.Get(&list, query, userId, listId)
	return list, err
}
func (r *NotesListPostgres) Update(userId, listId int, list model.ListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if list.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *list.Title)
		argId++
	}

	if list.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *list.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(
		`UPDATE %s list_table
		SET %s
		FROM %s list_user_table
		WHERE list_table.id = list_user_table.list_id AND
			list_user_table.user_id=$%d AND
			list_user_table.list_id=$%d`,
		noteListTable, setQuery, usersListTable, argId, argId+1)

	args = append(args, userId, listId)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)

	return err
}
func (r *NotesListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf(`
	DELETE FROM %s list_table
	USING %s list_user_table
	WHERE list_table.id = list_user_table.list_id AND
		list_user_table.user_id=$1 AND
		list_user_table.list_id=$2`,
		noteListTable, usersListTable)
	_, err := r.db.Exec(query, userId, listId)
	return err
}
