package api_v1

import (
	"app/api/api_v1/endpoints"
	"github.com/gin-gonic/gin"
)

func Api(router *gin.Engine) {
	router.GET("/", endpoints.GetHelloWorld)
}

