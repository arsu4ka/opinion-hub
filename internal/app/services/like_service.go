package services

import (
	"github.com/aru4ka/opinion-hub/internal/app/models"
	"github.com/aru4ka/opinion-hub/internal/app/repositories"
	"github.com/google/uuid"
)

type LikeService struct {
	repo repositories.ILikeRepository
}

func NewLikeService(repo repositories.ILikeRepository) *LikeService {
	return &LikeService{repo: repo}
}

func (ls *LikeService) Create(userId uint, opinionId uuid.UUID) error {
	like := &models.Like{
		ID:        uuid.New(),
		UserID:    userId,
		OpinionID: opinionId,
	}
	return ls.repo.Create(like)
}

func (ls *LikeService) Delete(id uuid.UUID) error {
	return ls.repo.Delete(id)
}

func (ls *LikeService) GetOpinionLikes(opinionId uuid.UUID) (int64, error) {
	return ls.repo.GetOpinionLikes(opinionId)
}
