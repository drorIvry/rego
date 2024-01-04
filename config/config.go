package config

import (
	"os"
	"strconv"

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

var TASK_TIMEOUT int
var SERVER_PORT int
var IN_CLUSTER bool

func InitConfig() {
	var err error
	DB_DRIVER = os.Getenv("DB_DRIVER")
	DB_SQLITE_URL = os.Getenv("DB_SQLITE_URL")

	DB_POSTGRES_HOST = os.Getenv("DB_POSTGRES_HOST")
	DB_POSTGRES_USERNAME = os.Getenv("DB_POSTGRES_USERNAME")
	DB_POSTGRES_PASSWORD = os.Getenv("DB_POSTGRES_PASSWORD")
	DB_POSTGRES_DB_NAME = os.Getenv("DB_POSTGRES_DB_NAME")
	DB_POSTGRES_DSN_EXTRA = os.Getenv("DB_POSTGRES_DSN_EXTRA")
	DB_POSTGRES_PORT, err = strconv.Atoi(os.Getenv("DB_POSTGRES_PORT"))
	if err != nil {
		DB_POSTGRES_PORT = 5432
		log.Println("Can't parse DB_POSTGRES_PORT, using default")
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
