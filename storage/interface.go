package storage

import "context"

//go:generate mockgen -source=interface.go -destination=mock_storage/mock.go -package=mock_storage

type StorageClient interface {
	GetFile(ctx context.Context, path string) ([]byte, error)
}
