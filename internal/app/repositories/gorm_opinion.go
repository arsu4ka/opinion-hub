package repositories

import (
	"github.com/aru4ka/opinion-hub/internal/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormOpinionRepository struct {
	db *gorm.DB
}

func NewOpinionRepository(db *gorm.DB) *GormOpinionRepository {
	return &GormOpinionRepository{
		db: db,
	}
}

func (os *GormOpinionRepository) FindById(id uuid.UUID) (*models.Opinion, error) {
	opinion := models.Opinion{ID: id}
	result := os.db.First(&opinion)
	return &opinion, result.Error
}

func (os *GormOpinionRepository) FindByUserID(userID uuid.UUID) ([]*models.Opinion, error) {
	var opinions []*models.Opinion
	result := os.db.Where("user_id = ?", userID).Find(&opinions)
	return opinions, result.Error
}

func (os *GormOpinionRepository) Create(opinion *models.Opinion) error {
	return os.db.Create(opinion).Error
}

func (os *GormOpinionRepository) Update(opinion *models.Opinion) error {
	return os.db.Save(opinion).Error
}

func (os *GormOpinionRepository) Delete(id uuid.UUID) error {
	return os.db.Delete(&models.Opinion{ID: id}).Error
}
