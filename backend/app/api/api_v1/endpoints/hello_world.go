package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func GetHelloWorld(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello World!")
}
