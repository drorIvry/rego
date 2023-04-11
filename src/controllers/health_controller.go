package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping				              godoc
// @Summary      				  Health check route
// @Description                   Used for health check
// @Tags                          Health
// @Produce                       json
// @Success                       200
// @Router                        /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
