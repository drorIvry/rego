package dao

import (
	"crypto/sha256"

	"github.com/rs/zerolog/log"

	"github.com/drorivry/rego/initializers"
	"github.com/drorivry/rego/models"
)

func CreateApiKey(apiKey *models.ApiKeys) error {
	result := initializers.GetTaskDefinitionsTable().Create(apiKey)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Error saving to database")
		return result.Error
	}
	return nil
}

func AuthApiKey(apiKey string) (models.ApiKeys, error) {
	var api_key models.ApiKeys

	h := sha256.New()
    h.Write([]byte(apiKey))
    hashedKey := string(h.Sum(nil))

	result := initializers.GetTaskDefinitionsTable().Where(
		"api_key = ?",
		hashedKey,
	).Find(&api_key)
	if result.Error != nil {
		log.Error().Err(result.Error).Msg("Error querying database")
		return api_key, result.Error
	}
	return api_key, nil
}