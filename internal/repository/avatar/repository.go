package avatar

import (
	"context"
	"fmt"

	"github.com/escoutdoor/bookit/internal/client/s3"
	"github.com/escoutdoor/bookit/internal/model"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type repository struct {
	cl s3.MinIOClient
}

func NewAvatarRepository(cl s3.MinIOClient) *repository {
	return &repository{
		cl: cl,
	}
}

func (r *repository) Upload(ctx context.Context, in *model.UploadAvatar) (string, error) {
	const op = "AvatarRepository.Upload"

	name := fmt.Sprintf("avatars/%s", uuid.New())
	err := r.cl.PutObject(
		ctx,
		name,
		in.Payload,
		in.Size,
		minio.PutObjectOptions{
			ContentType: "image/webp",
		},
	)
	if err != nil {
		return "", fmt.Errorf("%s: s3 put object: %s", op, err)
	}

	return r.cl.PublicURL(name), nil
}
