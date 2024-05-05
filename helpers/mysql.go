package helpers

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var MySQLClient *gorm.DB

const (
	//mysql: mysql://AINewsV2_shareblow:8e755072b72d0d61e51ae2205622134dd881b040@fhp.h.filess.io:3307/AINewsV2_shareblow
	tidb_user     = "AINewsV2_shareblow"
	tidb_password = "8e755072b72d0d61e51ae2205622134dd881b040"
	tidb_host     = "fhp.h.filess.io"
	tidb_port     = "3307"
	tidb_db_name  = "AINewsV2_shareblow"
)

func MysqlReconnect() (error, *gorm.DB) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
		tidb_user, tidb_password, tidb_host, tidb_port, tidb_db_name)
	MySQLClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		BuzzLogger.Warn(fmt.Sprintf("mysql connect error: %s \n", err.Error()))
		return err, nil
	}
	return nil, MySQLClient
}
func InitMysql() {

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
		tidb_user, tidb_password, tidb_host, tidb_port, tidb_db_name)
	MySQLClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Printf("mysql connect error: %s \n", err.Error())
		panic("mysql connect error: " + err.Error())
	}
	db, err := MySQLClient.DB()
	if err != nil {
		fmt.Printf("mysql get db error: %s \n", err.Error())
		panic("mysql get db error: " + err.Error())
	}
	db.SetConnMaxLifetime(time.Hour)
	fmt.Println("mysql connect success")
}
