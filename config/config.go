package config

import (
	"os"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/drorivry/rego/initializers"
)

var DB_URL string
var TASK_TIMEOUT int
var SERVER_PORT string
var IN_CLUSTER bool

func InitConfig() {
	initializers.LoadEnvVars()
	DB_URL = os.Getenv("DB_URL")
	parsedTimeout, err := strconv.Atoi(os.Getenv("TASK_TIMEOUT"))

	if err != nil {
		TASK_TIMEOUT = 300
		log.Error().Err(err).Msg("Can't parse tasktimeout, using default")
	}

	TASK_TIMEOUT = parsedTimeout
	SERVER_PORT = os.Getenv("SERVER_PORT")
	IN_CLUSTER, err = strconv.ParseBool(os.Getenv("IN_CLUSTER"))
	if err != nil {
		IN_CLUSTER = false
		log.Error().Err(err).Msg("Can't parse IN_CLUSTER, using default")
	}

}
