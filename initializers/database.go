package initializers

import (
	"fmt"
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
		sqlite.Open(config.DB_SQLITE_URL),
		&gorm.Config{},
	)
}

func connectPostgres() (*gorm.DB, error) {
	var dsn string
	if config.DB_POSTGRES_DSN != "" {
		dsn = config.DB_POSTGRES_DSN
	} else {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d %s",
			config.DB_POSTGRES_HOST,
			config.DB_POSTGRES_USERNAME,
			config.DB_POSTGRES_PASSWORD,
			config.DB_POSTGRES_DB_NAME,
			config.DB_POSTGRES_PORT,
			config.DB_POSTGRES_DSN_EXTRA,
		)
	}

	return gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
}

func connectMysql() (*gorm.DB, error) {
	var dsn string
	if config.DB_MYSQL_DSN != "" {
		dsn = config.DB_MYSQL_DSN
	} else {
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			config.DB_MYSQL_USERNAME,
			config.DB_MYSQL_PASSWORD,
			config.DB_MYSQL_HOST,
			config.DB_MYSQL_PORT,
			config.DB_MYSQL_DB_NAME,
		)
		if config.DB_MYSQL_DSN_EXTRA != "" {
			dsn += fmt.Sprintf(
				"?%s",
				config.DB_MYSQL_DSN_EXTRA,
			)
		}
	}

	return gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)
}

func migrateTables() {
	DB.AutoMigrate(&models.TaskDefinition{})
	DB.AutoMigrate(&models.TaskExecution{})
	DB.AutoMigrate(&models.ExecutionStatusHistory{})
}
