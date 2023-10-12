package services

import (
	"github.com/aru4ka/opinion-hub/internal/app/controllers/dto"
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

func (os *OpinionService) FindById(id uuid.UUID) (*models.Opinion, error) {
	return os.repo.FindById(id)
}

func (os *OpinionService) FindByUserId(userID uint, withDrafts bool) ([]*models.Opinion, error) {
	return os.repo.FindByUserId(userID, withDrafts)
}

func (os *OpinionService) Create(ownerId uint, opDto *dto.CreateOpinionDto) error {
	opinion := opDto.ToModel(ownerId)
	return os.repo.Create(opinion)
}

func (os *OpinionService) Update(id uuid.UUID, ownderId uint, opDto *dto.UpdateOpinionDto) error {
	opinion := opDto.ToModel(ownderId)
	opinion.ID = id
	return os.repo.Update(opinion)
}

func (os *OpinionService) Delete(id uuid.UUID) error {
	return os.repo.Delete(id)
}
