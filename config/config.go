package config

import (
	"log"
	"os"
	"strconv"

	"github.com/drorivry/matter/initializers"
)

var DB_URL string
var TASK_TIMEOUT int
var SERVER_PORT string

func InitConfig() {
	initializers.LoadEnvVars()
	DB_URL = os.Getenv("DB_URL")
	parsedTimeout, err := strconv.Atoi(os.Getenv("TASK_TIMEOUT"))

	if err != nil {
		TASK_TIMEOUT = 300
		log.Fatal("Can't parse tasktimeout, using default")
	}

	TASK_TIMEOUT = parsedTimeout
	SERVER_PORT = os.Getenv("SERVER_PORT")

}
