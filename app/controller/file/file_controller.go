package uploadcontroller

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yx1126/go-admin/common/minio"
	"github.com/yx1126/go-admin/response"
)

type FileController struct{}

const AvatarSize = 5 << 20 // 5MB

const AvatarBucketName = "go-admin"

func GenFileName(userId int, filename string) string {
	date := time.Now().UnixNano()
	suffix := strings.Split(filename, ".")
	return fmt.Sprintf("%d-%d-%s.%s", date, userId, uuid.New().String(), suffix[1])
}

func (*FileController) UploadAvatar(c *gin.Context) {
	file, _ := c.FormFile("file")
	if file == nil {
		response.NewError(nil).SetMsg("文件不能为空").Json(c)
		return
	}

	if file.Size > AvatarSize {
		response.NewError(nil).SetMsg("头像不能超过 5MB").Json(c)
		return
	}
	id := c.GetInt("userId")
	data, err := minio.UploadFile(AvatarBucketName, "avatar", id, file)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}
	response.NewSuccess(data).SetMsg("文件上传成功！").Json(c)
}

func (*FileController) GetFileObject(c *gin.Context) {
	filename := c.Query("filename")
	fmt.Println("filename:", filename)
	if filename == "" {
		response.NewError(nil).SetMsg("filename 不能为空").Json(c)
		return
	}
	fmt.Println("1")
	file, err := minio.GetFileObject(AvatarBucketName, filename)
	if err != nil {
		fmt.Println("2")
		response.NewError(err).Json(c)
		return
	}
	fmt.Println("3")
	defer file.Close()
	fmt.Println("4")

	stat, err := file.Stat()
	fmt.Println("stat", stat.Size)
	if err != nil {
		response.NewError(err).Json(c)
		return
	}

	c.Header("Content-Type", stat.ContentType)
	c.Header("Content-Disposition", "inline")
	c.Header("Content-Length", fmt.Sprintf("%d", stat.Size))

	if _, err = io.Copy(c.Writer, file); err != nil {
		response.NewError(err).Json(c)
	}
}
