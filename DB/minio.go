package DB

import (
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/yx1126/go-admin/config"
)

var Minio *minio.Client

func InitMinio() {
	address := config.Minio.Ip + ":" + strconv.Itoa(config.Minio.Port)
	minioClient, err := minio.New(address, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Minio.Access, config.Minio.Secret, ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
	Minio = minioClient
}
