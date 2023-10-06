package repositories

import (
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
	user := models.User{ID: id}
	result := us.db.First(&user)
	return &user, result.Error
}

func (us *GormUserRepository) FindByEmail(email string) (*models.User, error) {
	user := models.User{Email: email}
	result := us.db.First(&user)
	return &user, result.Error
}

func (us *GormUserRepository) FindByUsername(username string) (*models.User, error) {
	user := models.User{Username: username}
	result := us.db.First(&user)
	return &user, result.Error
}

func (us *GormUserRepository) Create(user *models.User) error {
	return us.db.Create(user).Error
}

func (us *GormUserRepository) Update(user *models.User) error {
	return us.db.Save(user).Error
}
