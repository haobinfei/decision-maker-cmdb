package aliyun

import (
	"decision-maker-cmdb/models"
	"fmt"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func (ac *AliCloud) SyncServerInfo() error {
	fmt.Println(ac.regionIds)
	for _, regionId := range ac.regionIds {
		request := &ecs20140526.DescribeInstancesRequest{
			RegionId:   tea.String(regionId),
			MaxResults: tea.Int32(ac.maxResults),
		}
		runtime := &util.RuntimeOptions{}
		for {
			resp, err := ac.client.DescribeInstancesWithOptions(request, runtime)
			if err != nil {
				fmt.Println(err)
				return err
			}
			for _, instancesWithOptions := range resp.Body.Instances.Instance {
				aliOSServe := &models.AssteServer{
					Hostname: *instancesWithOptions.HostName,
				}

				aliOSServe.Update()
			}
			if *resp.Body.NextToken == "" {
				break
			}
			request.NextToken = resp.Body.NextToken
		}

	}

	return nil

}
