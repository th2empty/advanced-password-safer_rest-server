package service

import (
	"github.com/th2empty/auth-server/pkg/models"
	"github.com/th2empty/auth-server/pkg/repository"
)

type PasswordListService struct {
	repo repository.PasswordList
}

func NewPasswordListService(repo repository.PasswordList) *PasswordListService {
	return &PasswordListService{repo: repo}
}

func (s *PasswordListService) Create(userId int, list models.PasswordList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *PasswordListService) GetAll(userId int) ([]models.PasswordList, error) {
	return s.repo.GetAll(userId)
}

func (s *PasswordListService) GetById(userId, id int) (models.PasswordList, error) {
	return s.repo.GetById(userId, id)
}
