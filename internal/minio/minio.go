package minio

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func Connect() *minio.Client {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalf("Error connecting to Minio: %v", err)
	}
	MinioClient = client
	fmt.Println("Successfully connected to Minio")
	return client
}

func MakeBucket(bucketName string) error {
	err := MinioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return err
	}
	return nil
}

func BucketExists(bucketName string) (bool, error) {
	found, err := MinioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		return false, err
	}
	return found, nil
}

func UploadImage(bucketName, objectName, contentType string, file multipart.File) error {
	_, err := MinioClient.PutObject(context.Background(), bucketName, objectName, file, -1, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return err
	}
	return nil
}

func DeleteImage(bucketName, objectName string) error {
	err := MinioClient.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
