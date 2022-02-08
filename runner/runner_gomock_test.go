package runner

import (
	"context"
	"errors"
	"testing"

	"github.com/cpurta/gomock-demo/storage/mock_storage"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRunnerNoError(t *testing.T) {
	ctrl := gomock.NewController(t)

	storageClient := mock_storage.NewMockStorageClient(ctrl)

	runner := NewRunner(storageClient)

	storageClient.EXPECT().GetFile(gomock.Any(), gomock.Any()).Return([]byte("Hello World"), nil).Times(3)

	err := runner.Run(context.Background())

	assert.NoError(t, err)
}

func TestRunnerError(t *testing.T) {
	ctrl := gomock.NewController(t)

	storageClient := mock_storage.NewMockStorageClient(ctrl)

	runner := NewRunner(storageClient)

	storageClient.EXPECT().GetFile(gomock.Any(), gomock.Any()).Return([]byte("Hello World"), nil)
	storageClient.EXPECT().GetFile(gomock.Any(), gomock.Any()).Return([]byte("Hello World"), nil)
	storageClient.EXPECT().GetFile(gomock.Any(), gomock.Any()).Return(nil, errors.New("forced error"))

	err := runner.Run(context.Background())

	assert.Error(t, err)
}
