package main

import (
	"log"

	"github.com/VinncentWong/Delos-AquaFarm/infrastructure"
)

func main() {
	logger := log.Default()

	// load environment variable
	err := infrastructure.LoadEnv()
	if err != nil {
		logger.Panic("error when load environment variable")
		return
	}

	// connect to database
	err = infrastructure.ConnectToDb()
	if err != nil {
		logger.Panic("error when trying to connect to db")
		return
	}

	// migrate
	err = infrastructure.Migrate()
	if err != nil {
		logger.Panic("error when migrate domains")
		return
	}

}
