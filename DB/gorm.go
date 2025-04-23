package db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/yx1126/go-admin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Gorm *gorm.DB

func InitGorm() {
	ms := config.Config.Mysql
	dsn := ms.User + ":" + ms.Password + "@tcp(" + ms.Ip + ":" + strconv.Itoa(ms.Port) + ")/" + ms.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("dsn: ", dsn)
	var err error
	Gorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := Gorm.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetConnMaxIdleTime(time.Duration(ms.MaxIdleConns))
	sqlDB.SetMaxOpenConns(ms.MaxOpenConns)
	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}
}
