package repositories

import (
	"errors"
	"github.com/aru4ka/opinion-hub/internal/app/models"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (us *GormUserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	result := us.db.Where("id = ?", id).First(&user)
	return &user, result.Error
}

func (us *GormUserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := us.db.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (us *GormUserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	result := us.db.Where("username = ?", username).First(&user)
	return &user, result.Error
}

func (us *GormUserRepository) Create(user *models.User) error {
	return us.db.Create(user).Error
}

func (us *GormUserRepository) Update(user *models.User) error {
	return us.db.Save(user).Error
}

func (us *GormUserRepository) CheckAvailability(username, email string) error {
	var user models.User
	us.db.Where("username = ? OR email = ?", username, email).First(&user)
	if user.Username == username {
		return errors.New("this username is already taken")
	} else if user.Email == email {
		return errors.New("this email is already taken")
	}
	return nil
}
