package models

import (
	"decision-maker-cmdb/conf"
	"log"
	"time"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var db *gorm.DB

func InitDB() {

	var err error
	db, err = gorm.Open(mysql.Open(conf.Config.GetString("mysql.username")+":"+conf.Config.GetString("mysql.password")+
		"@tcp("+conf.Config.GetString("mysql.host")+":"+conf.Config.GetString("mysql.port")+
		")/cmdb"+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		log.Fatalln("mysql连接失败")
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("mysql连接失败")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Minute * 60)

	db.AutoMigrate(&AssteServer{})

}
