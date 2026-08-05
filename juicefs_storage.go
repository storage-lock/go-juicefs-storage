package juicefs_storage

import (
	"context"

	s3_storage "github.com/storage-lock/go-s3-storage"
	"github.com/storage-lock/go-storage"
)

// JuicefsStorage juicefs对象存储实现 (S3兼容)
type JuicefsStorage struct {
	*s3_storage.S3Storage
}

func NewJuicefsStorage(ctx context.Context, options *s3_storage.S3StorageOptions) (*JuicefsStorage, error) {
	s3Storage, err := s3_storage.NewS3Storage(ctx, options)
	if err != nil {
		return nil, err
	}
	return &JuicefsStorage{S3Storage: s3Storage}, nil
}

const StorageName = "juicefs-storage"

func (x *JuicefsStorage) GetName() string {
	return StorageName
}

func (x *JuicefsStorage) Capabilities() []storage.StorageCapability {
	return x.S3Storage.Capabilities()
}
