package service

import (
	"github.com/th2empty/auth-server/pkg/models"
	"github.com/th2empty/auth-server/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type PasswordList interface {
	Create(userId int, list models.PasswordList) (int, error)
	GetAll(userId int) ([]models.PasswordList, error)
	GetById(userId, listId int) (models.PasswordList, error)
	Delete(userId, listId int) error
}

type PasswordItem interface {
	Add(userId, listId int, item models.PasswordItem) (int, error)
	GetAll(userId, listId int) ([]models.PasswordItem, error)
	GetById(userId, passwordId int) (models.PasswordItem, error)
	Delete(userId, passwordId int) error
	Update(userId, passwordId int, input models.UpdatePasswordInput) error
}

type Service struct {
	Authorization
	PasswordList
	PasswordItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		PasswordList:  NewPasswordListService(repos.PasswordList),
		PasswordItem:  NewPasswordItemService(repos.PasswordItem, repos.PasswordList),
	}
}
