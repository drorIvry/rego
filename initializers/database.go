package initializers

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/drorivry/rego/config"
	"github.com/drorivry/rego/models"
	"gorm.io/driver/mysql"
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

func GetApiKeysTable() *gorm.DB {
	return DB.Table(models.API_KEYS_TABLE_NAME)
}
func InitDBConnection() {
	var err error
	log.Info().Msg("initializing DB")

	if config.DB_DRIVER == "sqlite" {
		DB, err = connectSqlite()
	} else if config.DB_DRIVER == "postgresql" || config.DB_DRIVER == "postgres" {
		DB, err = connectPostgres()
	} else if config.DB_DRIVER == "mysql" {
		DB, err = connectMysql()
	} else {
		log.Error().Err(err).Str(
			"driver",
			config.DB_DRIVER,
		).Msg("DB Driver type is not supported")
		os.Exit(1)
	}

	if err != nil {
		log.Error().Err(err).Str(
			"driver",
			config.DB_DRIVER,
		).Msg("Error connecting to database")
		os.Exit(1)
	}

	migrateTables()
}

func connectSqlite() (*gorm.DB, error) {
	return gorm.Open(
		sqlite.Open(config.DB_URL),
		&gorm.Config{},
	)
}

func connectPostgres() (*gorm.DB, error) {
	return gorm.Open(
		postgres.Open(config.DB_URL),
		&gorm.Config{},
	)
}

func connectMysql() (*gorm.DB, error) {
	return gorm.Open(
		mysql.Open(config.DB_URL),
		&gorm.Config{},
	)
}

func migrateTable(model any, table_name string) {
	err := DB.AutoMigrate(model)
	if err != nil {
		log.Error().Err(err).Str(
			"table_name",
			table_name,
		).Msg(
			"Could not AutoMigrate table",
		)
	}
}

func migrateTables() {
	migrateTable(
		&models.TaskDefinition{},
		models.TASK_DEFINITIONS_TABLE_NAME,
	)
	migrateTable(
		&models.TaskExecution{},
		models.TASK_EXECUTIONS_TABLE_NAME,
	)
	migrateTable(
		&models.ExecutionStatusHistory{},
		models.EXECUTION_STATUS_HISTORY_TABLE_NAME,
	)
	migrateTable(
		&models.ApiKeys{},
		models.API_KEYS_TABLE_NAME,
	)
}
