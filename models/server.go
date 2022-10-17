package models

import (
	"time"
)

type AssteTag struct {
	ID       int64
	ServerId int64
	TagId    int64
}

type AssteDb struct {
	ID             int64
	Idc            int64  `gorm:"idc"`          // IDC编号
	DbInstanceId   string `gorm:"dbInstanceId"` // RDS实例ID
	DbInstanceName string `gorm:"dbInstanceName"`
	DbCode         string `gorm:"dbCode"`
	DbClass        string `gorm:"dbClass"`
	DbPublicIp     string `gorm:"dbPublicIp"`
	DbPort         string `gorm:"dbPort"`
	DbUser         string `gorm:"dbUser"`
	DbPwd          string `gorm:"dbPwd"`
	DbDisk         string `gorm:"dbDisk"`
	DbType         string `gorm:"dbType"`
	DbVersion      string `gorm:"dbVersion"`
	DbMark         string `gorm:"dbMark"`
	DbStatus       string `gorm:"dbStatus"`
	CreateAt       time.Time
	UpdateAt       time.Time
}

type AssteServer struct {
	ID        int64
	Hostname  string `gorm:"hostname"`
	Ip        string `gorm:"ip"`
	PublicIp  string `gorm:"publicIp"`
	PrivateIp string `gorm:"privateIp"`
	Port      int64  `gorm:"port"`
	Idc       int64  `gorm:"idc"` // IDC编号
	AdminUser string `gorm:"adminUser"`
	Region    string `gorm:"region"`
	State     string `gorm:"state"`
	Detail    string `gorm:"detail"`
	CreateAt  time.Time
	UpdateAt  time.Time
}

type ServerDetail struct {
	ID            int64
	Sn            string `gorm:"sn"`
	Ip            string `gorm:"ip"`
	Cpu           string `gorm:"cpu"`
	CpuCores      int64  `gorm:"cpuCores"`
	Memory        int64  `gorm:"memory"`
	Disk          int64  `gorm:"disk"`
	OsTpye        string `gorm:"osTpye"`
	OsKernel      string `gorm:"osKernel"`
	InstanceId    string `gorm:"instanceId"`
	InstanceType  string `gorm:"instanceType"`
	InstanceState string `gorm:"instanceState"`
	CreateAt      time.Time
	UpdateAt      time.Time
}
