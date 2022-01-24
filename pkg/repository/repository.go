package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/th2empty/auth-server/pkg/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type PasswordList interface {
	Create(userId int, list models.PasswordList) (int, error)
	GetAll(userId int) ([]models.PasswordList, error)
	GetById(userId, id int) (models.PasswordList, error)
}

type PasswordItem interface {
}

type Repository struct {
	Authorization
	PasswordList
	PasswordItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		PasswordList:  NewPasswordListPostgres(db),
	}
}
