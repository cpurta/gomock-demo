package runner

import (
	"context"
	"errors"
	"testing"

	"github.com/cpurta/gomock-demo/storage/my_mock"
	"github.com/stretchr/testify/assert"
)

func TestRunnerMyMockNoError(t *testing.T) {
	mockGetFile := func(ctx context.Context, path string) ([]byte, error) {
		return []byte("here is a testing string"), nil
	}

	mockStorage := my_mock.NewMockedStorageClient(mockGetFile)

	runner := NewRunner(mockStorage)

	err := runner.Run(context.Background())

	assert.NoError(t, err)
}

func TestRunnerMyMockError(t *testing.T) {
	mockGetFile := func(ctx context.Context, path string) ([]byte, error) {
		return nil, errors.New("forced error")
	}

	mockStorage := my_mock.NewMockedStorageClient(mockGetFile)

	runner := NewRunner(mockStorage)

	err := runner.Run(context.Background())

	assert.Error(t, err)
}
