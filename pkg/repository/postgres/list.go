package postgres

import (
	"fmt"
	"notes/model"

	"github.com/jmoiron/sqlx"
)

type NotesListPostgres struct {
	db *sqlx.DB
}

func NewNotesListPostgres(db *sqlx.DB) *NotesListPostgres {
	return &NotesListPostgres{db: db}
}

func (repository *NotesListPostgres) CreateList(userId int, list model.NotesList) (int, error) {
	transaction, err := repository.db.Begin()
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
func (repository *NotesListPostgres) GetAllListUserId(userId int) ([]model.NotesList, error) {
	var lists []model.NotesList

	query := fmt.Sprintf(
		`SELECT list_table.id, list_table.title, list_table.description 
		FROM %s list_table 
		INNER JOIN %s list_user_table on 
			list_table.id = list_user_table.list_id 
		WHERE list_user_table.user_id = $1`,
		noteListTable, usersListTable,
	)
	err := repository.db.Select(&lists, query, userId)
	return lists, err
}
func (repository *NotesListPostgres) GetListByIdUserId(userId, listId int) (model.NotesList, error) {
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
	err := repository.db.Get(&list, query, userId, listId)
	return list, err
}
func (repository *NotesListPostgres) UpdateListByIdUserId(userId, listId int, list model.UpdateListInput) error {
	return nil
}
func (repository *NotesListPostgres) DeleteListByIdUserId(userId, listId int) error {
	return nil
}
