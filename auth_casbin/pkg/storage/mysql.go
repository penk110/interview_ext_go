package storage

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/penk110/interview_ext_go/auth_casbin/conf"
)

var defaultGorm *gorm.DB

func Init(conf *conf.Conf) error {
	var err error
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "dev:Dev@3306@tcp(192.168.3.11:3306)/casbin?charset=utf8mb4&parseTime=True&loc=Local"
	defaultGorm, err = gorm.Open(mysql.Open(conf.Storage.DSN), &gorm.Config{})
	if err != nil {
		return err
	}
	mysqlDB, err := defaultGorm.DB()
	if err != nil {
		return err
	}
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)

	// open debug

	log.Println("DB.Ping(): ", mysqlDB.Ping())
	return nil
}

func DB() *gorm.DB {
	return defaultGorm
}
