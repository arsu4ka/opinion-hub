package repositories

import (
	"github.com/aru4ka/opinion-hub/internal/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormOpinionRepository struct {
	db *gorm.DB
}

func NewGormOpinionRepository(db *gorm.DB) *GormOpinionRepository {
	return &GormOpinionRepository{
		db: db,
	}
}

func (os *GormOpinionRepository) FindById(id uuid.UUID) (*models.Opinion, error) {
	var opinion models.Opinion
	result := os.db.Where("id = ?", id).First(&opinion)
	return &opinion, result.Error
}

func (os *GormOpinionRepository) FindByUserId(userID uint, withDrafts bool) ([]*models.Opinion, error) {
	var opinions []*models.Opinion
	query := "owner_id = ?"
	if !withDrafts {
		query = "owner_id = ? AND is_draft = false"
	}
	result := os.db.Where(query, userID).Find(&opinions)
	return opinions, result.Error
}

func (os *GormOpinionRepository) Create(opinion *models.Opinion) error {
	return os.db.Create(opinion).Error
}

func (os *GormOpinionRepository) Update(opinion *models.Opinion) error {
	return os.db.Save(opinion).Error
}

func (os *GormOpinionRepository) Delete(id uuid.UUID) error {
	return os.db.Delete(&models.Opinion{}, id).Error
}
