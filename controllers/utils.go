package controllers

import (
	"errors"

	"github.com/drorivry/rego/dao"
	"github.com/drorivry/rego/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func AuthRequest(c *gin.Context) (*models.ApiKeys, error){
	token := c.GetHeader("X-Rego-Api-Key")

	apiKey, authErr := dao.AuthApiKey(token)

	if authErr != nil {
		log.Error().Err(authErr).Msg("Invalid token")
		return nil, errors.New("Invalid token");
	}

	return &apiKey, nil

}