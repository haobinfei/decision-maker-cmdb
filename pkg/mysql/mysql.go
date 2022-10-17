package mysql

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	var db *sql.DB

	db, _ = sql.Open("mysql", "root:passwd@tcp(127.0.0.1:3306)/cmdb?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.SetConnMaxLifetime(time.Minute * 60)
	if err := db.Ping(); err != nil {
		log.Fatalln("mysql连接失败")
	}
	return db
}
