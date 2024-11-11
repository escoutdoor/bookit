package s3

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
)

type MinIOClient interface {
	PutObject(ctx context.Context, objectName string, payload io.Reader, size int64, opts minio.PutObjectOptions) error
	PublicURL(name string) string
}
