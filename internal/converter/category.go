package converter

import (
	"github.com/escoutdoor/bookit/internal/controller/http/dto"
	"github.com/escoutdoor/bookit/internal/model"
)

func ToCreateCategoryFromDTO(dto *dto.CreateCategoryDTO) *model.CreateCategory {
	return &model.CreateCategory{
		Name: dto.Name,
	}
}

func ToUpdateCategoryFromDTO(dto *dto.UpdateCategoryDTO) *model.UpdateCategory {
	return &model.UpdateCategory{
		Name: dto.Name,
	}
}
