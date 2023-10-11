package services

import (
	"github.com/aru4ka/opinion-hub/internal/app/models"
	"github.com/aru4ka/opinion-hub/internal/app/repositories"
	"github.com/google/uuid"
)

type OpinionService struct {
	repo repositories.IOpinionRepository
}

func NewOpinionService(repo repositories.IOpinionRepository) *OpinionService {
	return &OpinionService{repo: repo}
}

func (os *OpinionService) FindByID(id uuid.UUID) (*models.Opinion, error) {
	return os.repo.FindByID(id)
}

func (os *OpinionService) FindByUserID(userID uint, withDrafts bool) ([]*models.Opinion, error) {
	return os.repo.FindByUserID(userID, withDrafts)
}

func (os *OpinionService) Create(opinion *models.Opinion) error {
	if err := opinion.Validate(); err != nil {
		return err
	}

	return os.repo.Create(opinion)
}

func (os *OpinionService) Update(opinion *models.Opinion) error {
	if err := opinion.Validate(); err != nil {
		return err
	}

	return os.repo.Update(opinion)
}

func (os *OpinionService) Delete(id uuid.UUID) error {
	return os.repo.Delete(id)
}
