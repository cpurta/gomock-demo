package my_mock

import (
	"context"

	"github.com/cpurta/gomock-demo/storage"
)

var _ storage.StorageClient = &mockedStorageClient{}

type mockedStorageClient struct {
	getFileFunc func(ctx context.Context, path string) ([]byte, error)
}

func NewMockedStorageClient(f func(ctx context.Context, path string) ([]byte, error)) *mockedStorageClient {
	return &mockedStorageClient{
		getFileFunc: f,
	}
}

func (mocked *mockedStorageClient) GetFile(ctx context.Context, path string) ([]byte, error) {
	return mocked.getFileFunc(ctx, path)
}
