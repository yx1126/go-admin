package DB

import (
	"strconv"

	"go-admin/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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
