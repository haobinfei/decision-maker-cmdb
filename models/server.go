package models

import (
	"log"
	"time"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v3/client"
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
	ID          uint64 `gorm:"primaryKey"`
	InstanceId  string `gorm:"primaryKey"`
	Hostname    string
	Ip          string
	PublicIp    string
	PrivateIp   string
	Idc         int64
	AdminUser   string
	Region      string
	State       string
	Description string
	CreateAt    time.Time
	UpdateAt    time.Time
}

type ServerDetail struct {
	ID            int64
	Sn            string
	Ip            string
	Cpu           string
	CpuCores      int64
	Memory        int64
	Disk          int64
	OsTpye        string
	OsKernel      string
	InstanceId    string
	InstanceType  string
	InstanceState string
	CreateAt      time.Time
	UpdateAt      time.Time
}

func update(assteServer *AssteServer) {
	result := db.Save(assteServer)
	if result.Error != nil {
		log.Fatalln("主机信息更新失败")
	}
}

func create(assteServer *AssteServer) {
	result := db.Omit("CreateAt", "UpdateAt").Create(assteServer)
	if result.Error != nil {
		log.Fatalln("主机信息插入失败")
	}
}

func isExist(instanceId string) bool {
	var server []AssteServer
	db.Where("instance_id = ?", instanceId).Find(&server)
	log.Println(server)
	return len(server) != 0
}

func Sync(DescribeInstances *ecs20140526.DescribeInstancesResponseBodyInstancesInstance) {
	var privateIp string
	var publicIp string
	if *DescribeInstances.InstanceNetworkType == "vpc" {
		privateIp = *DescribeInstances.VpcAttributes.PrivateIpAddress.IpAddress[0]
	} else {
		privateIp = *DescribeInstances.InnerIpAddress.IpAddress[0]
	}
	if len(DescribeInstances.PublicIpAddress.IpAddress) != 0 {
		publicIp = *DescribeInstances.PublicIpAddress.IpAddress[0]
	}

	instanceId := *DescribeInstances.InstanceId

	assteServer := &AssteServer{
		InstanceId:  instanceId,
		Hostname:    *DescribeInstances.HostName,
		Ip:          *DescribeInstances.EipAddress.IpAddress,
		PublicIp:    publicIp,
		Region:      *DescribeInstances.RegionId,
		State:       *DescribeInstances.Status,
		Description: *DescribeInstances.Description,
		PrivateIp:   privateIp,
	}
	if isExist(instanceId) {
		update(assteServer)
	} else {
		create(assteServer)
	}

}
