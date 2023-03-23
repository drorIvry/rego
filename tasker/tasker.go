package tasker

import (
	"log"
	"net/http"

	"github.com/drorivry/matter/controllers"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		log.Fatal("Error", err)
	}

	c.JSON(http.StatusInternalServerError, "")
}

func GetServer() *gin.Engine {
	r := gin.Default()
	r.Use(ErrorHandler)

	r.GET("/ping", controllers.Ping)
	r.POST("/task", controllers.CreateTaskDefinition)
	r.GET("/task", controllers.GetAllTaskDefinitions)
	r.GET("/tasks/pending", controllers.GetAllPendingTaskDefinitions)

	return r
}
