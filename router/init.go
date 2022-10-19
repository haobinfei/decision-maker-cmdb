package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.New()
	loadAsset(r)
	return r
}
