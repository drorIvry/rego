package tasker

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}
