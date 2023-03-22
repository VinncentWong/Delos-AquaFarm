package main

import (
	"fmt"
	"log"
	"os"

	farmHandler "github.com/VinncentWong/Delos-AquaFarm/app/farm/handler"
	farmRepository "github.com/VinncentWong/Delos-AquaFarm/app/farm/repository"
	farmUsecase "github.com/VinncentWong/Delos-AquaFarm/app/farm/usecase"
	pondHandler "github.com/VinncentWong/Delos-AquaFarm/app/pond/handler"
	pondRepository "github.com/VinncentWong/Delos-AquaFarm/app/pond/repository"
	pondUsecase "github.com/VinncentWong/Delos-AquaFarm/app/pond/usecase"
	"github.com/VinncentWong/Delos-AquaFarm/infrastructure"
	"github.com/VinncentWong/Delos-AquaFarm/rest"
	"github.com/gin-gonic/gin"
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

	// init router
	router := gin.Default()

	// init farm
	fRepo := farmRepository.NewFarmRepository()
	fUsecase := farmUsecase.NewFarmUsecase(fRepo)
	fHandler := farmHandler.NewFarmHandler(fUsecase)

	// init pond
	pRepo := pondRepository.NewPondRepository()
	pUsecase := pondUsecase.NewPondUsecase(pRepo, fRepo)
	pHandler := pondHandler.NewPondHandler(pUsecase)

	// init routing on rest directory
	routing := rest.NewRouting(router)
	routing.InitializeCheckHealthRouting()
	routing.InitializeFarmRouting(fHandler)
	routing.InitializePondRouting(pHandler)

	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
