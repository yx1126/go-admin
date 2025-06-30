package uploadcontroller

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yx1126/go-admin/common/minio"
	"github.com/yx1126/go-admin/response"
)

type UploadController struct{}

const AvatarSize = 5 << 20 // 5MB

const AvatarBucketName = "go-admin"

func GenFileName(userId int, filename string) string {
	date := time.Now().UnixNano()
	suffix := strings.Split(filename, ".")
	return fmt.Sprintf("%d-%d-%s.%s", date, userId, uuid.New().String(), suffix[1])
}

func (*UploadController) UploadAvatar(c *gin.Context) {
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
