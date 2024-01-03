package config

import (
	"log"
	"os"
	"strconv"

	"github.com/drorivry/rego/initializers"
)

var DB_URL string
var TASK_TIMEOUT int
var SERVER_PORT int
var IN_CLUSTER bool

func InitConfig() {
	var err error
	initializers.LoadEnvVars()
	DB_URL = os.Getenv("DB_URL")

	TASK_TIMEOUT, err = strconv.Atoi(os.Getenv("TASK_TIMEOUT"))
	if err != nil {
		TASK_TIMEOUT = 300
		log.Println("Can't parse tasktimeout, using default")
	}

	SERVER_PORT, err = strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		SERVER_PORT = 4004
		log.Println("Can't parse port, using default")
	}

	IN_CLUSTER, err = strconv.ParseBool(os.Getenv("IN_CLUSTER"))
	if err != nil {
		IN_CLUSTER = false
		log.Println("Can't parse IN_CLUSTER, using default")
	}

}
