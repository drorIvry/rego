package config

import (
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

var DB_DRIVER string
var DB_URL string

var TASK_TIMEOUT int
var SERVER_PORT int
var IN_CLUSTER bool

func InitConfig() {
	var err error
	DB_DRIVER = strings.ToLower(os.Getenv("DB_DRIVER"))
	if !slices.Contains(
		[]string{
			"sqlite",
			"postgres",
			"postgresql",
			"mysql",
		},
		DB_DRIVER,
	) {
		log.Error().Msg("Not supported db driver")
		os.Exit(1)
	}
	DB_URL = os.Getenv("DB_URL")

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
