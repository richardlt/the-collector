package files

import (
	"bytes"
	"fmt"
	"io"

	minio "github.com/minio/minio-go"
	"github.com/pkg/errors"
	"github.com/richardlt/the-collector/server/types"
)

var minioClient struct {
	client *minio.Client
	bucket string
}

var minioBucket string

// InitStorage .
func InitStorage(minioURI, minioAccessKey, minioSecretKey,
	minioBucket string, minioSSL bool) error {
	client, err := minio.New(minioURI, minioAccessKey, minioSecretKey,
		minioSSL)
	if err != nil {
		return errors.WithStack(err)
	}
	minioClient.client = client
	minioClient.bucket = minioBucket

	exists, err := minioClient.client.BucketExists(minioBucket)
	if err != nil {
		return errors.WithStack(err)
	}
	if !exists {
		return errors.WithStack(minioClient.client.MakeBucket(minioBucket, ""))
	}

	return nil
}

// ReadImage .
func ReadImage(path string, size string) ([]byte, error) {
	if size != types.Original {
		path = fmt.Sprintf("%s.%s", path, size)
	}

	o, err := minioClient.client.GetObject(minioClient.bucket, path,
		minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, o); err != nil {
		if v, ok := err.(minio.ErrorResponse); ok && v.Code == "NoSuchKey" {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}

	return buf.Bytes(), nil
}

// SaveFile .
func SaveFile(data []byte, path string) error {
	r := bytes.NewReader(data)
	_, err := minioClient.client.PutObject(minioClient.bucket, path, r,
		r.Size(), minio.PutObjectOptions{})
	return errors.WithStack(err)
}
