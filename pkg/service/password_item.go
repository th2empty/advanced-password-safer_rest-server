package service

import (
	"github.com/th2empty/auth-server/pkg/models"
	"github.com/th2empty/auth-server/pkg/repository"
)

type PasswordItemService struct {
	repo     repository.PasswordItem
	listRepo repository.PasswordList
}

func NewPasswordItemService(repo repository.PasswordItem, listRepo repository.PasswordList) *PasswordItemService {
	return &PasswordItemService{repo: repo, listRepo: listRepo}
}

func (s *PasswordItemService) Add(userId, listId int, item models.PasswordItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil { // list does not exists or does not belongs to user
		return 0, err
	}

	return s.repo.Add(listId, item)
}

func (s *PasswordItemService) GetAll(userId, listId int) ([]models.PasswordItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *PasswordItemService) GetById(userId, passwordId int) (models.PasswordItem, error) {
	return s.repo.GetById(userId, passwordId)
}

func (s *PasswordItemService) Delete(userId, passwordId int) error {
	return s.repo.Delete(userId, passwordId)
}

func (s *PasswordItemService) Update(userId, passwordId int, input models.UpdatePasswordInput) error {
	return s.repo.Update(userId, passwordId, input)
}
