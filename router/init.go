package router

import "github.com/gin-gonic/gin"

func InitHttpServe() *gin.Engine {
	r := gin.New()
	loadAsset(r)
	return r
}
