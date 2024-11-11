package avatar

import (
	"context"
	"fmt"
	"io"

	"github.com/escoutdoor/bookit/internal/model"
)

func (uc *usecase) Upload(ctx context.Context, payload io.Reader, size int64) (string, error) {
	const op = "AvatarUseCase.Upload"

	in := &model.UploadAvatar{
		Payload: payload,
		Size:    size,
	}
	url, err := uc.repo.Upload(ctx, in)
	if err != nil {
		return "", fmt.Errorf("%s: %s", op, err)
	}

	return url, nil
}
