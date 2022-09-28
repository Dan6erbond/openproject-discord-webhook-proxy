package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/flytam/filenamify"
	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func init() {
	viper.SetDefault("storage.s3.endpoint", "s3.amazonaws.com")
	viper.SetDefault("storage.s3.usessl", true)
}

func GetFilename(path string) (string, error) {
	t := time.Now()
	filename, err := filenamify.Filenamify(fmt.Sprintf(
		"%d-%02d-%02dT%02d:%02d:%02d_%s.json",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), path),
		filenamify.Options{Replacement: "_"})
	if err != nil {
		return "", err
	}
	return filename, nil
}

type StorageService interface {
	SaveRequest(body []byte, path string) error
}

type NilStorageService struct {
}

func (nss *NilStorageService) SaveRequest(body []byte, path string) error {
	return nil
}

type LocalStorageService struct {
	path   string
	logger *log.Logger
}

func (lss *LocalStorageService) SaveRequest(body []byte, path string) error {
	filename, err := GetFilename(path)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(viper.GetString("storage.local.path"), filename), body, 0644)
	return err
}

type S3StorageService struct {
	minioClient *minio.Client
	bucketName  string
	logger      *log.Logger
}

func (s3 *S3StorageService) SaveRequest(body []byte, path string) error {
	ctx := context.Background()
	filename, err := GetFilename(path)
	if err != nil {
		return err
	}
	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		os.Mkdir("tmp", 0755)
	}
	filepath := filepath.Join("tmp", filename)
	err = os.WriteFile(filepath, body, 0644)
	if err != nil {
		return err
	}
	contentType := "application/json"
	objectName := "/requests/" + filename
	info, err := s3.minioClient.FPutObject(ctx, s3.bucketName, objectName, filepath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}
	s3.logger.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	err = os.Remove(filepath)
	if err != nil {
		return err
	}
	return nil
}

func NewS3StorageService(logger *log.Logger, secure bool, bucketName, region, endpoint, accessKey, secretKey string) *S3StorageService {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: secure,
	})
	if err != nil {
		logger.Fatalln(err)
	}

	return &S3StorageService{minioClient, bucketName, logger}
}

func NewStorageService(lc fx.Lifecycle, logger *log.Logger, router *mux.Router) StorageService {
	logger.Print("Executing NewStorageService.")
	s3Config := viper.GetStringMapString("storage.s3")
	var storageService StorageService
	if s3Bucket, ok := s3Config["bucketname"]; ok {
		storageService = NewS3StorageService(
			logger,
			viper.GetBool("storage.s3.usessl"),
			s3Bucket,
			s3Config["region"],
			s3Config["endpoint"],
			s3Config["accesskey"],
			s3Config["secretkey"],
		)
		return storageService
	}

	if viper.GetString("storage.local.path") != "" {
		storageService = &LocalStorageService{path: viper.GetString("storage.local.path"), logger: logger}
		return storageService
	}

	storageService = &NilStorageService{}
	return storageService
}
