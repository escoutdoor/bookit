package minio

import (
	"context"
	"fmt"
	"github.com/escoutdoor/bookit/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewConn(ctx context.Context, cfg *config.S3Config) (*minio.Client, error) {
	const op = "minio.NewConn"
	cl, err := minio.New(cfg.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AcessKey, cfg.SecretKey, ""),
		Region: cfg.Region,
		Secure: *cfg.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("%s: connect to s3 storage: %s", op, err)
	}

	ie, err := cl.BucketExists(ctx, cfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("%s: check bucket exists: %s", op, err)
	}
	if !ie {
		if err := cl.MakeBucket(ctx, cfg.Bucket, minio.MakeBucketOptions{}); err != nil {
			return nil, fmt.Errorf("%s: make bucket: %s", op, err)
		}
		if err := cl.SetBucketPolicy(
			ctx,
			cfg.Bucket,
			`{
				"Version": "2012-10-17",
				"Statement": [
					{
						"Effect": "Allow",
						"Principal": {
							"AWS": [
								"*"
							]
						},
						"Action": [
							"s3:GetObject"
						],
						"Resource": [
							"arn:aws:s3:::`+cfg.Bucket+`/*"
						]
					}
				]
			}`); err != nil {
			return nil, fmt.Errorf("%s: error SetBucketPolicy: %s", op, err)
		}
	}

	return cl, err
}
