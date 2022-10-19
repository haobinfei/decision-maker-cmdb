package models

import (
	"decision-maker-cmdb/conf"
	"log"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func InitDB() {

	var err error
	db, err = gorm.Open("mysql", conf.Config.GetString("mysql.username")+":"+conf.Config.GetString("mysql.password")+
		"@tcp("+conf.Config.GetString("mysql.host")+":"+conf.Config.GetString("mysql.port")+
		")/cmdb"+"?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalln("mysql连接失败")
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Minute * 60)

	db.SingularTable(true)
	db.AutoMigrate(&AssteServer{})

}
