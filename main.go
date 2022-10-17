package main

import (
	"decision-maker-cmdb/conf"
	"decision-maker-cmdb/pkg/mysql"
)

func main() {

	conf.InitConfig()
	db := mysql.InitDB()
	db.Ping()

}
