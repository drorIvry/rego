package initializers

import (
	"log"

	"github.com/drorivry/matter/models"
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
	log.Println("initializing DB")
	var err error
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&models.TaskDefinition{})
	DB.AutoMigrate(&models.TaskExecution{})

}
