package services

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

type StorageService interface {
	SaveRequest()
}

type LocalStorageService struct {
	path string
}

func (lss *LocalStorageService) SaveRequest() {
}

type S3StorageService struct {
	minioClient *minio.Client
	bucketName  string
}

func (s3 *S3StorageService) SaveRequest() {
	ctx := context.Background()
	objectName := uuid.NewString()
	filePath := "/requests/" + objectName
	contentType := "application/json"
	info, err := s3.minioClient.FPutObject(ctx, s3.bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}

func NewS3StorageService(bucketName, region, endpoint, accessKey, secretKey string) *S3StorageService {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return &S3StorageService{minioClient, bucketName}
}

func NewStorageService() StorageService {
	s3Config := viper.GetStringMapString("storage.s3")
	if s3Bucket, ok := s3Config["bucketName"]; ok {
		s3StorageService := NewS3StorageService(
			s3Bucket,
			s3Config["region"],
			s3Config["endpoint"],
			s3Config["accessKey"],
			s3Config["secretKey"],
		)
		return s3StorageService
	}

	localStorageService := LocalStorageService{path: viper.GetString("storage.local.path")}
	return &localStorageService
}
