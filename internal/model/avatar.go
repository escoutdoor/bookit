package model

import "io"

type UploadAvatar struct {
	Payload io.Reader
	Size    int64
}
