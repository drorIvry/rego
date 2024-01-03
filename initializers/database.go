package initializers

import (
	"github.com/rs/zerolog/log"

	"github.com/drorivry/rego/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetTaskDefinitionsTable() *gorm.DB {
	return DB.Table("task_definitions")
}

func GetTaskExecutionsTable() *gorm.DB {
	return DB.Table("task_executions")
}

func InitDBConnection(dbName string) {
	log.Info().Msg("initializing DB")
	var err error
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&models.TaskDefinition{})
	DB.AutoMigrate(&models.TaskExecution{})
}
