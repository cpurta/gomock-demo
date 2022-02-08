package s3_client

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cpurta/gomock-demo/storage"
)

var _ storage.StorageClient = &s3Client{}

type s3Client struct {
	service *s3.S3
	bucket  string
}

func NewS3Client(service *s3.S3, bucket string) *s3Client {
	return &s3Client{
		service: service,
		bucket:  bucket,
	}
}

func (client *s3Client) GetFile(ctx context.Context, path string) ([]byte, error) {
	input := &s3.GetObjectInput{
		Bucket: &client.bucket,
		Key:    &path,
	}

	output, err := client.service.GetObject(input)
	if err != nil {
		return nil, fmt.Errorf("Unable to get object from s3: %s", err.Error())
	}

	fileBytes, err := ioutil.ReadAll(output.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to read object body: %s", err.Error())
	}

	return fileBytes, nil
}
