package gorm

import (
	"fmt"

	"TimeLine/internal/config"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB mysql数据库engine配置，在main开始时初始化
var (
	DB *gorm.DB
)

func Init() {
	var err error
	mysqlHost := config.MustString("mysql_ip")
	mysqlUser := config.MustString("mysql_user")
	mysqlPassword := config.MustString("mysql_password")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/cucc-sdwan?charset=utf8mb4", mysqlUser, mysqlPassword, mysqlHost)

	// 初始化mysql engine
	DB, err = gorm.Open(mysql.Open(dataSource))
	if err != nil {
		panic(err)
	}

	log.Info("Success to connect mysql")
}
