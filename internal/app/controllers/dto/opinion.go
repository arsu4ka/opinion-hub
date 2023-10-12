package dto

import (
	"github.com/aru4ka/opinion-hub/internal/app/models"
	"github.com/google/uuid"
)

type CreateOpinionDto struct {
	Title   string `json:"title" xml:"title" validate:"required,min=3,max=255"`
	Body    string `json:"body" xml:"body"`
	IsDraft bool   `json:"isDraft" xml:"isDraft" validate:"required"`
}

func (od *CreateOpinionDto) ToModel(ownerId uint) *models.Opinion {
	return &models.Opinion{
		Title:   od.Title,
		Body:    od.Body,
		IsDraft: od.IsDraft,
		OwnerID: ownerId,
	}
}

type UpdateOpinionDto struct {
	CreateOpinionDto
}

type ResponseOpinionDto struct {
	ID      uuid.UUID `json:"id" xml:"id"`
	Title   string    `json:"title" xml:"title"`
	Body    string    `json:"body" xml:"body"`
	IsDraft bool      `json:"isDraft" xml:"isDraft"`
	Likes   uint      `json:"likes" xml:"likes"`
	OwnerID uint      `json:"ownerId" xml:"ownerId"`
}

func NewResponseOpinionDto(opinion *models.Opinion) *ResponseOpinionDto {
	return &ResponseOpinionDto{
		ID:      opinion.ID,
		Title:   opinion.Title,
		Body:    opinion.Body,
		IsDraft: opinion.IsDraft,
		Likes:   opinion.Likes,
		OwnerID: opinion.OwnerID,
	}
}
