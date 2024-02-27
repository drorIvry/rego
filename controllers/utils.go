package controllers

import (
	"errors"
	"strconv"

	"github.com/drorivry/rego/dao"
	"github.com/drorivry/rego/models"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func AuthRequest(c *gin.Context) (*models.ApiKeys, error) {
	token := c.GetHeader("X-Rego-Api-Key")

	apiKey, authErr := dao.AuthApiKey(token)

	if authErr != nil {
		log.Error().Err(authErr).Msg("Invalid token")
		return nil, errors.New("Invalid token")
	}

	return &apiKey, nil

}

func ParseIntQueryParameter(c *gin.Context, paramName string, defaultValue int) int {
	param_str := c.DefaultQuery(paramName, strconv.Itoa(defaultValue))
	param, err := strconv.Atoi(param_str)

	if err != nil {
		param = 0
	}

	return param
}
