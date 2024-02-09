package dao

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/rs/zerolog/log"

	"github.com/drorivry/rego/initializers"
	"github.com/drorivry/rego/models"
)

func CreateApiKey(apiKey *models.ApiKeys) error {
	result := initializers.GetApiKeysTable().Create(apiKey)
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
	hashedKey := hex.EncodeToString(h.Sum(nil))

	result := initializers.GetApiKeysTable().Where(
		"api_key = ?",
		hashedKey,
	).First(&api_key)
	if result.Error != nil || result.RowsAffected == 0 {
		log.Error().Err(result.Error).Msg("Error querying database")
		return api_key, errors.New("Invalid token")
	}
	return api_key, nil
}