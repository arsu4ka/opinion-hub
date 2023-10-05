package services

import (
	"github.com/aru4ka/opinion-hub/internal/app/models"
	"github.com/aru4ka/opinion-hub/internal/app/repositories"
)

type UserService struct {
	repo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) FindById(id uint) (*models.User, error) {
	return us.repo.FindByID(id)
}

func (us *UserService) FindByEmail(email string) (*models.User, error) {
	return us.repo.FindByEmail(email)
}

func (us *UserService) FindByUsername(username string) (*models.User, error) {
	return us.repo.FindByUsername(username)
}

func (us *UserService) Create(user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	return us.repo.Create(user)
}

func (us *UserService) Update(user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	return us.repo.Update(user)
}
