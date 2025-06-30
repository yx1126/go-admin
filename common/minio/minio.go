package minio

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/yx1126/go-admin/DB"
)

func GenFileName(userId int, filename string) string {
	date := time.Now().UnixNano()
	suffix := strings.Split(filename, ".")
	return fmt.Sprintf("%d-%d-%s.%s", date, userId, uuid.New().String(), suffix[1])
}

func UploadFile(bucket, folder string, userId int, fileHeader *multipart.FileHeader) (any, error) {
	if fileHeader == nil {
		return nil, errors.New("文件不能为空！")
	}
	// 自动创建 bucket
	exists, err := DB.Minio.BucketExists(context.Background(), bucket)
	if err != nil {
		return nil, err
	}
	if !exists {
		if err := DB.Minio.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{}); err != nil {
			return nil, err
		}
	}
	file, err := fileHeader.Open()
	defer file.Close()
	if err != nil {
		return nil, err
	}
	filename := GenFileName(userId, fileHeader.Filename)
	path := fmt.Sprintf("/%s/%s", folder, filename)
	info, err := DB.Minio.PutObject(context.Background(), bucket, path, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: fileHeader.Header.Get("Content-Type"),
	})
	if err != nil {
		return nil, err
	}
	return gin.H{
		"fileName":     filename,
		"size":         info.Size,
		"lastModified": info.LastModified.Local().Format("2006-01-02 15:04:05"),
		"path":         fmt.Sprintf("/%s%s", bucket, path),
	}, nil
}
