package main

import (
	"context"
	"fmt"

	"github.com/cpurta/gomock-demo/runner"
	"github.com/cpurta/gomock-demo/storage/filesystem"
)

func main() {
	fsClient := filesystem.NewFileSystemClient()

	runner := runner.NewRunner(fsClient)

	if err := runner.Run(context.Background()); err != nil {
		fmt.Println("Error running runner:", err)
	}
}
