package category

import (
	"context"
	"errors"
	"fmt"

	categoryrepo "github.com/escoutdoor/bookit/internal/repository/category"
	"github.com/google/uuid"
)

func (uc *usecase) Delete(ctx context.Context, id uuid.UUID) error {
	const op = "CategoryUseCase.Delete"

	err := uc.repo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, categoryrepo.ErrNotFound) {
			return ErrNotFound
		}

		return fmt.Errorf("%s: %s", op, err)
	}

	return nil
}
