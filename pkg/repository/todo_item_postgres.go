package repository

import (
	"fmt"
	"webapp/model"

	"github.com/jmoiron/sqlx"
)

type todoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *todoItemPostgres {
	return &todoItemPostgres{db: db}
}

// in process...
func (r *todoItemPostgres) Create(userId, listId int, item model.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	
	var itemId int

	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, desctription) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := r.db.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemQuery := fmt.Sprintf("INSERT INTO %s (listId, itemId) VALUES($1, $2)", listsItemsTable)
	_, err = r.db.Exec(createListItemQuery, listId, itemId)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

// in process...
func (r *todoItemPostgres) GetAll(userId, listId int) ([]model.TodoItem, error) {
	var items []model.TodoItem

	query := fmt.Sprintf("Select ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s tl on tl.item_id = ti.id INNER JOIN %s ul on ul.list_id = tl.list_id WHERE tl.list_d = $1 AND ul.user_id = $2", 
							todoItemsTable, listsItemsTable, usersListsTable)

	err := r.db.Select(&items, query, listId, userId)

	return items, err 
}

// in process...
func (r *todoItemPostgres) GetById(userId, itemId int) (model.TodoItem, error) {
	var item model.TodoItem

	query := fmt.Sprintf("Select ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on ti.id = li.item_id INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2", 
							todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}
	
	return item, nil
}

// in process...
func (r *todoItemPostgres) Delete(userId, listId, itemId int) error {
	return nil
}
