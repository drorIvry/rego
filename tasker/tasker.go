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

	c.JSON(http.StatusInternalServerError, "Internal Server Error")
}

func GetServer() *gin.Engine {
	r := gin.Default()
	r.Use(ErrorHandler)

	r.GET("/ping", controllers.Ping) // V
	r.POST("/task", controllers.CreateTaskDefinition) // V
	r.GET("/task", controllers.GetAllTaskDefinitions) // V
	r.POST("/task/:definitionId/rerun", controllers.RerunTask) // dror
	r.GET("/task/:definitionId/latest") // dror
	r.PUT("/task") // dror
	r.DELETE("/task/:definitionId") // dror

	r.GET("/execution", controllers.RerunTask)
	r.POST("/execution/:execId/abort", controllers.RerunTask)
	r.GET("/tasks/pending", controllers.GetAllPendingTaskDefinitions)

	return r
}
