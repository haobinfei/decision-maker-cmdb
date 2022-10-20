package aliyun

// ali SDK verson is 2014-05-26

import (
	"decision-maker-cmdb/conf"
	"time"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

type AliCloud struct {
	accessKeyId     *string
	accessKeySecret *string
	client          *ecs20140526.Client
	regionIds       []string
	maxResults      int32
	ChSync          chan bool
	ChExit          chan bool
}

var AliOpt *AliCloud

func NewAliCloud() {
	AliOpt = new(AliCloud)
	AliOpt.accessKeyId = tea.String(conf.Config.GetString("aliyun.accessKeyId"))
	AliOpt.accessKeySecret = tea.String(conf.Config.GetString("aliyun.accessKeySecret"))
	AliOpt.regionIds = conf.Config.GetStringSlice("aliyun.regionId")
	AliOpt.maxResults = 10
	AliOpt.client = AliOpt.createClient()
	AliOpt.ChSync = make(chan bool)
	AliOpt.ChExit = make(chan bool)

}

func (ac *AliCloud) createClient() *ecs20140526.Client {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: ac.accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: ac.accessKeySecret,
	}
	config.Endpoint = tea.String("ecs.cn-wulanchabu.aliyuncs.com")
	client, _ := ecs20140526.NewClient(config)
	return client
}

func (ac *AliCloud) StartSyncOS() {
	syncTicker := time.NewTicker(2 * time.Hour)
	for {
		select {
		case <-ac.ChSync:
			syncTicker.Reset(2 * time.Hour)
			ac.SyncServerInfo()
		case <-syncTicker.C:
			ac.SyncServerInfo()
		case <-ac.ChExit:
			return
		}
	}
}

func (ac *AliCloud) StartSync() {
	ac.ChSync <- true
}
