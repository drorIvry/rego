package initializers

import (
	"log"

	"github.com/drorivry/matter/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DefinitionsTable *gorm.DB
var ExecutionsTable *gorm.DB

func InitDBConnection(dbName string) {
	log.Println("initializing DB")
	var err error
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&models.TaskDefinition{})
	DB.AutoMigrate(&models.TaskExecution{})

	// Create Table References
	DefinitionsTable = DB.Table("task_definitions")
	ExecutionsTable = DB.Table("task_executions")
}
