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
	_, authErr := AuthRequest(c)

	if authErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}


// Health				              godoc
// @Summary      				  Health check route
// @Description                   Used for health check
// @Tags                          Health
// @Produce                       json
// @Success                       200
// @Router                        / [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "healthy",
	})
}
