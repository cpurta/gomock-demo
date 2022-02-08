package filesystem

import (
	"context"
	"os"

	"github.com/cpurta/gomock-demo/storage"
)

var _ storage.StorageClient = &fileSystemClient{}

type fileSystemClient struct{}

func NewFileSystemClient() *fileSystemClient {
	return &fileSystemClient{}
}

func (fs *fileSystemClient) GetFile(ctx context.Context, path string) ([]byte, error) {
	return os.ReadFile(path)
}
