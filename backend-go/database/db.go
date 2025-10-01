package database

import (
	"backend-go/config"
	"backend-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() error {
	dsn := config.AppConfig.MySQL.GetDSN()
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 自动迁移数据表
	err = DB.AutoMigrate(&models.User{}, &models.Paper{}, &models.Question{})
	if err != nil {
		log.Printf("数据库迁移警告: %v", err)
	}

	log.Println("数据库连接成功")
	return nil
} 