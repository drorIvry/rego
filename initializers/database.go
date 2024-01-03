package initializers

import (
	"log"

	"github.com/drorivry/rego/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetTaskDefinitionsTable() *gorm.DB {
	return DB.Table(models.TASK_DEFINITIONS_TABLE_NAME)
}

func GetTaskExecutionsTable() *gorm.DB {
	return DB.Table(models.TASK_EXECUTIONS_TABLE_NAME)
}

func GetExecutionsStatusHistoryTable() *gorm.DB {
	return DB.Table(models.EXECUTION_STATUS_HISTORY_TABLE_NAME)
}

func InitDBConnection(dbName string) {
	log.Println("initializing DB " + dbName)
	var err error
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Migrate the schema
	DB.AutoMigrate(&models.TaskDefinition{})
	DB.AutoMigrate(&models.TaskExecution{})
	DB.AutoMigrate(&models.ExecutionStatusHistory{})
}
