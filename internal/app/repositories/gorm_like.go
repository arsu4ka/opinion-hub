package repositories

import (
	"github.com/aru4ka/opinion-hub/internal/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormLikeRepository struct {
	db *gorm.DB
}

func NewGormLikeRepository(db *gorm.DB) *GormLikeRepository {
	return &GormLikeRepository{db: db}
}

func (r *GormLikeRepository) Create(like *models.Like) error {
	return r.db.Create(like).Error
}

func (r *GormLikeRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Like{ID: id}).Error
}

func (r *GormLikeRepository) GetOpinionLikes(opinionId uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&models.Like{}).Where("opinion_id = ?", opinionId).Count(&count).Error
	return count, err
}
