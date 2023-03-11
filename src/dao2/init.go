package dao2

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"winter-examination/src/model"
)

var db *gorm.DB

func init() {
	db = DBInit()
}

func GetDB() *gorm.DB {
	return db
}

func DBInit() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	dsn := "root:root@tcp(127.0.0.1:3306)/winter_examination_database?charset=utf8mb4&parseTime=True&loc=Local"
	newDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil
	}
	err = newDB.AutoMigrate(&model.Client{})
	if err != nil {
		fmt.Printf("err = %v", err)
		return nil
	}
	return newDB
}
