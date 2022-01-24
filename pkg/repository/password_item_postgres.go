package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/th2empty/auth-server/pkg/models"
	"strings"
)

type PasswordItemPostgres struct {
	db *sqlx.DB
}

func NewPasswordItemPostgres(db *sqlx.DB) *PasswordItemPostgres {
	return &PasswordItemPostgres{db: db}
}

func (r *PasswordItemPostgres) Add(listId int, item models.PasswordItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "repository",
			"file":     "password_item_postgres.go",
			"function": "Add",
			"message":  err,
		}).Errorf("error while starting transaction")
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf(`INSERT INTO %s (title, web_site, login, phone, email, pass, secret_word, backup_codes) 
										VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`, passwordsItemTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.WebSite, item.Login, item.Phone, item.Email, item.Pass, item.SecretWord, item.BackupCodes)
	if err := row.Scan(&itemId); err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "repository",
			"file":     "password_item_postgres.go",
			"function": "Add",
			"message":  err,
		}).Errorf("scan scopies returned error")
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf(`INSERT INTO %s (list_id, item_id) values ($1, $2)`, listItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "repository",
			"file":     "password_item_postgres.go",
			"function": "Add",
			"message":  err,
		}).Errorf("error while execute a query")
		tx.Rollback()
		return 0, err
	}

	logrus.WithFields(logrus.Fields{
		"package":  "repository",
		"file":     "password_item_postgres.go",
		"function": "Add",
	}).Info("transaction successfully done")

	return itemId, tx.Commit()
}

func (r *PasswordItemPostgres) GetAll(userId, listId int) ([]models.PasswordItem, error) {
	var items []models.PasswordItem
	query := fmt.Sprintf(`SELECT pi.id, pi.title, pi.web_site, pi.login, pi.phone, pi.email, pi.pass, pi.secret_word, pi.backup_codes FROM %s pi
								INNER JOIN %s li on li.item_id = pi.id 
								INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		passwordsItemTable, listItemsTable, usersListTable)

	if err := r.db.Select(&items, query, listId, userId); err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "repository",
			"file":     "password_item_postgres.go",
			"function": "GetAll",
			"message":  err,
		}).Errorf("error when using select from database")
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"package":  "repository",
		"file":     "password_item_postgres.go",
		"function": "GetAll",
	}).Infof("query successfully done")

	return items, nil
}

func (r *PasswordItemPostgres) GetById(userId, passwordId int) (models.PasswordItem, error) {
	var password models.PasswordItem
	query := fmt.Sprintf(`SELECT pi.id, pi.title, pi.web_site, pi.login, pi.phone, pi.email, pi.pass, pi.secret_word, pi.backup_codes FROM %s pi
								INNER JOIN %s li on li.item_id = pi.id 
								INNER JOIN %s ul on ul.list_id = li.list_id WHERE pi.id = $1 AND ul.user_id = $2`,
		passwordsItemTable, listItemsTable, usersListTable)

	if err := r.db.Get(&password, query, passwordId, userId); err != nil {
		logrus.WithFields(logrus.Fields{
			"package":  "repository",
			"file":     "password_item_postgres.go",
			"function": "GetById",
			"message":  err,
		}).Errorf("error while getting data from db")
		return password, err
	}

	logrus.WithFields(logrus.Fields{
		"package":  "repository",
		"file":     "password_item_postgres.go",
		"function": "GetById",
	}).Infof("query successfully done")

	return password, nil
}

func (r *PasswordItemPostgres) Delete(userId, passwordId int) error {
	query := fmt.Sprintf(`DELETE FROM %s pi USING %s li, %s ul WHERE pi.id = li.item_id AND li.list_id = ul.list_id
								AND ul.user_id = $1 AND pi.id = $2`,
		passwordsItemTable, listItemsTable, usersListTable)

	_, err := r.db.Exec(query, userId, passwordId)

	return err
}

func (r *PasswordItemPostgres) Update(userId, passwordId int, input models.UpdatePasswordInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.WebSite != nil {
		setValues = append(setValues, fmt.Sprintf("web_site=$%d", argId))
		args = append(args, *input.WebSite)
		argId++
	}

	if input.Login != nil {
		setValues = append(setValues, fmt.Sprintf("login=$%d", argId))
		args = append(args, *input.Login)
		argId++
	}

	if input.Phone != nil {
		setValues = append(setValues, fmt.Sprintf("phone=$%d", argId))
		args = append(args, *input.Phone)
		argId++
	}

	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *input.Email)
		argId++
	}

	if input.Pass != nil {
		setValues = append(setValues, fmt.Sprintf("pass=$%d", argId))
		args = append(args, *input.Pass)
		argId++
	}

	if input.SecretWord != nil {
		setValues = append(setValues, fmt.Sprintf("secret_word=$%d", argId))
		args = append(args, *input.SecretWord)
		argId++
	}

	if input.BackupCodes != nil {
		setValues = append(setValues, fmt.Sprintf("backup_codes=$%d", argId))
		args = append(args, *input.BackupCodes)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s pi SET %s FROM %s li, %s ul 
								WHERE pi.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d 
								AND pi.id = $%d`,
		passwordsItemTable, setQuery, listItemsTable, usersListTable, argId, argId+1)
	args = append(args, userId, passwordId)

	_, err := r.db.Exec(query, args...)
	return err
}
