package minio

import (
	"log"
	"work01/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func NewMinioClient() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("can not connect to config")
	}

	cfg := config.ReadInConfig()

	endpoint := "localhost:9000"
	accessKeyID := cfg.MINIO_ACCESS_KEY
	secretAccessKey := cfg.MINIO_SECRET_KEY
	useSSL := false

	var err error
	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("can not connect to MinIO: %v", err)
	}

	log.Println("Connect to MinIO Success")
}
