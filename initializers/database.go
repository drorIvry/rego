package initializers

import (
	"log"
	"time"

	models "github.com/drorivry/matter/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDBConnection(dbName string) {
	log.Println("initializing DB")
	var err error
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	args := []string{"aa"}

	taskDef := models.TaskDefinition{
		Image:                   "hello-world",
		TtlSecondsAfterFinished: 10,
		Status:                  models.READY,
		ExecutionInterval:       10,
		ExecutionsCounter:       0,
		NextExecutionTime:       time.Now(),
		Enabled:                 true,
		Deleted:                 false,
		Args:                    args,
		Cmd:                     "",
	}

	// Migrate the schema
	DB.AutoMigrate(&models.TaskDefinition{})
	DB.AutoMigrate(&models.TaskExecution{})

	DB.Create(&taskDef)

}
