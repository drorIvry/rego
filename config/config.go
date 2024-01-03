package config

import (
	"log"
	"os"
	"strconv"
)

var DB_DRIVER string
var DB_SQLITE_URL string
var DB_POSTGRES_DSN string

var TASK_TIMEOUT int
var SERVER_PORT string
var IN_CLUSTER bool

func InitConfig() {
	DB_DRIVER = os.Getenv("DB_DRIVER")
	DB_SQLITE_URL = os.Getenv("DB_SQLITE_URL")
	DB_POSTGRES_DSN = os.Getenv("DB_POSTGRES_DSN")

	parsedTimeout, err := strconv.Atoi(os.Getenv("TASK_TIMEOUT"))
	if err != nil {
		TASK_TIMEOUT = 300
		log.Fatal("Can't parse tasktimeout, using default")
	}

	TASK_TIMEOUT = parsedTimeout
	SERVER_PORT = os.Getenv("SERVER_PORT")
	IN_CLUSTER, err = strconv.ParseBool(os.Getenv("IN_CLUSTER"))
	if err != nil {
		IN_CLUSTER = false
		log.Fatal("Can't parse IN_CLUSTER, using default")
	}
}
