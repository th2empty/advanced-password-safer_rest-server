package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/th2empty/auth-server/pkg/models"
)

type PasswordListPostgres struct {
	db *sqlx.DB
}

func NewPasswordListPostgres(db *sqlx.DB) *PasswordListPostgres {
	return &PasswordListPostgres{db: db}
}

func (r *PasswordListPostgres) Create(userId int, list models.PasswordList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", passwordListTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES($1, $2)", usersListTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *PasswordListPostgres) GetAll(userId int) ([]models.PasswordList, error) {
	var lists []models.PasswordList

	query := fmt.Sprintf("SELECT pl.id, pl.title, pl.description FROM %s pl INNER JOIN %s ul on pl.id = ul.list_id WHERE ul.user_id = $1",
		passwordListTable, usersListTable)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *PasswordListPostgres) GetById(userId, listId int) (models.PasswordList, error) {
	var list models.PasswordList

	query := fmt.Sprintf(`SELECT pl.id, pl.title, pl.description FROM %s pl 
								INNER JOIN %s ul on pl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
		passwordListTable, usersListTable)

	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *PasswordListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s pl USING %s ul WHERE pl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2",
		passwordListTable, usersListTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}
