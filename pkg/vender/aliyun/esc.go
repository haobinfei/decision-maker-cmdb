package aliyun

// func init() {
// 	router.SyncList["aliyun"] = HanderSyncOSServer
// }

// func HanderSyncOSServer(regionIds []string) error {
// 	for _, regionId := range regionIds {
// 		AliCloudEcsUpdate(regionId)
// 	}
// 	return nil
// }

// func AliCloudEcsUpdate(regionId string) {

// }

import (
	"decision-maker-cmdb/conf"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AliCloud struct {
	AccessKeyId     *string
	AccessKeySecret *string
	client          *ecs.Client
	regionId        []string
	maxResults      int32
}

func NewAliCloud() *AliCloud {
	aliCloud := new(AliCloud)
	aliCloud.AccessKeyId = tea.String(conf.Config.GetString("AccessKeyId"))
	aliCloud.AccessKeySecret = tea.String(conf.Config.GetString("AccessKeySecret"))
	aliCloud.regionId = conf.Config.GetStringSlice("regionId")
	aliCloud.client = aliCloud.createClient()
	return aliCloud
}

func (ac *AliCloud) createClient() *ecs.Client {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: ac.AccessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: ac.AccessKeySecret,
	}
	config.Endpoint = tea.String("ecs.beijing-cn.aliyuncs.com")
	client, _ := ecs.NewClient(config)
	return client
}

func (ac *AliCloud) GetServerInfo() {
	request := &ecs.DescribeInstancesRequest{
		RegionId:   tea.String(ac.regionId),
		MaxResults: tea.Int32(ac.maxResults),
	}
	runtime := &util.RuntimeOptions{}

	for {
		resp, err := ac.client.DescribeInstancesWithOptions(request, runtime)
		if err != nil {
			return
		}

		if resp.Body.NextToken == nil {
			break
		}
		request.NextToken = resp.Body.NextToken
	}
}
