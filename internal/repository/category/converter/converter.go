package converter

import (
	"github.com/escoutdoor/bookit/internal/model"
	repomodel "github.com/escoutdoor/bookit/internal/repository/category/model"
)

func ToCategoryFromRepository(c *repomodel.Category) *model.Category {
	return &model.Category{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
	}
}
