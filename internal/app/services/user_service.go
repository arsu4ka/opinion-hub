package services

import (
	"github.com/aru4ka/opinion-hub/internal/app/controllers/dto"
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
	return us.repo.FindById(id)
}

func (us *UserService) FindByEmail(email string) (*models.User, error) {
	return us.repo.FindByEmail(email)
}

func (us *UserService) FindByUsername(username string) (*models.User, error) {
	return us.repo.FindByUsername(username)
}

func (us *UserService) Create(userDto *dto.CreateUserDto) error {
	user := userDto.ToModel()
	return us.repo.Create(user)
}

func (us *UserService) Update(id uint, userDto *dto.UpdateUserDto) error {
	user := userDto.ToModel()
	user.ID = id
	return us.repo.Update(user)
}

func (us *UserService) CheckAvailability(username, email string) error {
	return us.repo.CheckAvailability(username, email)
}
