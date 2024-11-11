package minio

import (
	"context"
	"fmt"
	"io"

	"github.com/escoutdoor/bookit/internal/client/s3"
	"github.com/escoutdoor/bookit/internal/config"
	"github.com/minio/minio-go/v7"
)

type client struct {
	minioClient *minio.Client
	cfg         *config.S3Config
}

var _ s3.MinIOClient = (*client)(nil)

func NewClient(ctx context.Context, cfg *config.S3Config) (*client, error) {
	const op = "minio.NewClient"
	minioConn, err := NewConn(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: new minio connection: %s", op, err)
	}

	return &client{
		minioClient: minioConn,
		cfg:         cfg,
	}, nil
}

func (cl *client) PutObject(
	ctx context.Context,
	objectName string,
	payload io.Reader,
	size int64,
	opts minio.PutObjectOptions,
) error {
	const op = "MinioClient.PutObject"
	_, err := cl.minioClient.PutObject(
		ctx,
		cl.cfg.Bucket,
		objectName,
		payload,
		size,
		opts,
	)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	return nil
}

func (cl *client) PublicURL(id string) string {
	return fmt.Sprintf("%s/%s", cl.cfg.PublicURL, id)
}
