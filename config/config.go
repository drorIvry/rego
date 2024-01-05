package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

var DB_DRIVER string
var DB_SQLITE_URL string

var DB_POSTGRES_HOST string
var DB_POSTGRES_PORT int
var DB_POSTGRES_USERNAME string
var DB_POSTGRES_PASSWORD string
var DB_POSTGRES_DB_NAME string
var DB_POSTGRES_DSN_EXTRA string

var DB_MYSQL_USERNAME string
var DB_MYSQL_PASSWORD string
var DB_MYSQL_HOST string
var DB_MYSQL_PORT int
var DB_MYSQL_DB_NAME string
var DB_MYSQL_DSN_EXTRA string

var TASK_TIMEOUT int
var SERVER_PORT int
var IN_CLUSTER bool

func initPostgresConfig() {
	var err error
	DB_POSTGRES_HOST = os.Getenv("DB_POSTGRES_HOST")
	DB_POSTGRES_USERNAME = os.Getenv("DB_POSTGRES_USERNAME")
	DB_POSTGRES_PASSWORD = os.Getenv("DB_POSTGRES_PASSWORD")
	DB_POSTGRES_DB_NAME = os.Getenv("DB_POSTGRES_DB_NAME")
	DB_POSTGRES_DSN_EXTRA = os.Getenv("DB_POSTGRES_DSN_EXTRA")
	DB_POSTGRES_PORT, err = strconv.Atoi(os.Getenv("DB_POSTGRES_PORT"))
	if err != nil {
		DB_POSTGRES_PORT = 5432
		log.Error().Err(err).Msg("Can't parse DB_POSTGRES_PORT, using default")
	}
}

func initMySqlConfig() {
	var err error
	DB_MYSQL_USERNAME = os.Getenv("DB_MYSQL_USERNAME")
	DB_MYSQL_PASSWORD = os.Getenv("DB_MYSQL_PASSWORD")
	DB_MYSQL_HOST = os.Getenv("DB_MYSQL_HOST")
	DB_MYSQL_DB_NAME = os.Getenv("DB_MYSQL_DB_NAME")
	DB_MYSQL_DSN_EXTRA = os.Getenv("DB_MYSQL_DSN_EXTRA")
	DB_MYSQL_PORT, err = strconv.Atoi(os.Getenv("DB_MYSQL_PORT"))
	if err != nil {
		DB_POSTGRES_PORT = 3306
		log.Error().Err(err).Msg("Can't parse DB_MYSQL_PORT, using default")
	}
}

func initSqliteConfig() {
	DB_SQLITE_URL = os.Getenv("DB_SQLITE_URL")
}

func InitConfig() {
	var err error
	DB_DRIVER = strings.ToLower(os.Getenv("DB_DRIVER"))

	if DB_DRIVER == "postgres" {
		initPostgresConfig()
	} else if DB_DRIVER == "mysql" {
		initMySqlConfig()
	} else if DB_DRIVER == "sqlite" {
		initSqliteConfig()
	} else {
		log.Error().Msg("Not supported db driver")
		os.Exit(1)
	}

	TASK_TIMEOUT, err = strconv.Atoi(os.Getenv("TASK_TIMEOUT"))
	if err != nil {
		TASK_TIMEOUT = 300
		log.Error().Err(err).Msg("Can't parse TASK_TIMEOUT, using default")
	}

	SERVER_PORT, err = strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		SERVER_PORT = 4004
		log.Error().Err(err).Msg("Can't parse SERVER_PORT, using default")
	}

	IN_CLUSTER, err = strconv.ParseBool(os.Getenv("IN_CLUSTER"))
	if err != nil {
		IN_CLUSTER = false
		log.Error().Err(err).Msg("Can't parse IN_CLUSTER, using default")
	}
}
