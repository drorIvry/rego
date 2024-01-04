package initializers

import (
	"log"

	"github.com/drorivry/rego/config"
	"github.com/drorivry/rego/models"
	"gorm.io/driver/postgres"
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

func InitDBConnection() {
	var err error
	log.Println("initializing DB")

	if config.DB_DRIVER == "sqlite" {
		DB, err = connectSqlite()
	} else if config.DB_DRIVER == "postgresql" || config.DB_DRIVER == "postgres" {
		DB, err = connectPostgres()
	} else {
		log.Fatal("DB Driver type is not supported " + config.DB_DRIVER)
	}

	if err != nil {
		log.Fatal("Error connecting to database")
		return
	}

	migrateTables()
}

func connectSqlite() (*gorm.DB, error) {
	return gorm.Open(
		sqlite.Open(config.DB_SQLITE_URL),
		&gorm.Config{},
	)
}

func connectPostgres() (*gorm.DB, error) {
	return gorm.Open(
		postgres.Open(config.DB_POSTGRES_DSN),
		&gorm.Config{},
	)
}

func migrateTables() {
	DB.AutoMigrate(&models.TaskDefinition{})
	DB.AutoMigrate(&models.TaskExecution{})
	DB.AutoMigrate(&models.ExecutionStatusHistory{})
}
