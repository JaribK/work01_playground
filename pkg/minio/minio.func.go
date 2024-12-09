package minio

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"path"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func UploadAvatar(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	defer file.Close()

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	if err != nil {
		return "", err
	}

	extension := fileHeader.Filename[strings.LastIndex(fileHeader.Filename, "."):]
	ojbName := fmt.Sprintf("%s%d%s", strings.Replace(uuid.New().String(), "-", "", -1), time.Now().UnixNano(), extension)
	bucketName := "testlocal"

	ctx := context.Background()
	_, err = MinioClient.PutObject(ctx, bucketName, ojbName, buffer, int64(buffer.Len()), minio.PutObjectOptions{
		ContentType: fileHeader.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", err
	}

	fileURL := fmt.Sprintf("%s/%s/%s", MinioClient.EndpointURL().String(), bucketName, ojbName)
	return fileURL, nil
}

func UploadAvatarUpdate(fileHeader *multipart.FileHeader, avatartURL string) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}

	defer file.Close()

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	if err != nil {
		return "", err
	}

	bucketName := "testlocal"
	ctx := context.Background()

	if avatartURL != "" {
		Fullpath := avatartURL
		oldOjbName := path.Base(Fullpath)

		err = MinioClient.RemoveObject(ctx, bucketName, oldOjbName, minio.RemoveObjectOptions{})
		if err != nil {
			log.Printf("error removing old file: %v", err)
		} else {
			fmt.Printf("Old file %s deleted success", oldOjbName)
		}
	}

	extension := fileHeader.Filename[strings.LastIndex(fileHeader.Filename, "."):]
	ojbName := fmt.Sprintf("%s%d%s", strings.Replace(uuid.New().String(), "-", "", -1), time.Now().UnixNano(), extension)

	_, err = MinioClient.PutObject(ctx, bucketName, ojbName, buffer, int64(buffer.Len()), minio.PutObjectOptions{
		ContentType: fileHeader.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", err
	}

	fileURL := fmt.Sprintf("%s/%s/%s", MinioClient.EndpointURL().String(), bucketName, ojbName)
	return fileURL, nil
}
