package DB

import (
	"strconv"

	config "github.com/yx1126/go-admin/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Gorm *gorm.DB

func InitGorm() {
	ms := config.Config.Mysql
	dsn := ms.User + ":" + ms.Password + "@tcp(" + ms.Ip + strconv.Itoa(ms.Port) + ")/" + ms.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Gorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
