package router

import (
	"github.com/gin-gonic/gin"
)

func loadAsset(r *gin.Engine) {
	cmdb := r.Group("/api/cmdb")
	cmdb.GET("asset_config/handler_update_server", HanderUpdateOSServer)
	cmdb.GET("asset_config/handler_update_rds", HanderUpdateRDS)
	cmdb.GET("asset_config/handler_update_redis", HanderUpdateRedis)
}

// var SyncList = make(map[string]func([]string) error)

func HanderUpdateOSServer(c *gin.Context) {

}

func HanderUpdateRDS(c *gin.Context) {

}

func HanderUpdateRedis(c *gin.Context) {

}
