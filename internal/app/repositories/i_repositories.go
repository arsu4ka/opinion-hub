package repositories

import (
	"github.com/aru4ka/opinion-hub/internal/app/models"
	"github.com/google/uuid"
)

type IUserRepository interface {
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	CheckAvailability(username, email string) error
}

type IOpinionRepository interface {
	FindByID(id uuid.UUID) (*models.Opinion, error)
	FindByUserID(userID uint, withDrafts bool) ([]*models.Opinion, error)
	Create(opinion *models.Opinion) error
	Update(opinion *models.Opinion) error
	Delete(id uuid.UUID) error
}

type ILikeRepository interface {
	Create(like *models.Like) error
	Delete(id uuid.UUID) error
	GetOpinionLikes(opinionId uuid.UUID) (int64, error)
}
