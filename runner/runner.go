package runner

import (
	"context"
	"fmt"

	"github.com/cpurta/gomock-demo/storage"
)

type Runner struct {
	storageClient storage.StorageClient
}

func NewRunner(storageClient storage.StorageClient) *Runner {
	return &Runner{
		storageClient: storageClient,
	}
}

func (runner *Runner) Run(ctx context.Context) error {
	for i := 0; i < 3; i++ {
		filePath := fmt.Sprintf("testing/example_%d.txt", i)

		fileBytes, err := runner.storageClient.GetFile(ctx, filePath)
		if err != nil {
			return fmt.Errorf("Unable to read file %s: %s", filePath, err.Error())
		}

		fmt.Println(string(fileBytes))
	}

	return nil
}
