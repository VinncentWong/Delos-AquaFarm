# Tech stack
- Go v1.20
- [Gin Gonic Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/) as ORM written in Go
- [Postgresql](https://www.postgresql.org/) as RDBMS
- [Testify](https://github.com/stretchr/testify) as a toolkit for assertion, mocking, etc.
# Installation
1. Clone this repository
2. Type on your terminal `cd Delos-AquaFarm/`
3. Make a `.env` file
4. Copy all key from `.env.example` file to `.env` file. Fill all the blank value with your credential 
5. Turn on the storage by typing this command on terminal :  `docker-compose -f docker-storage.yaml up -d`
6. To run all test on the application, type this command on terminal : `GIN_MODE=release go test ./... -v`
7. To run the application. type this command on terminal : `go run main.go`
# Additional information
All API documentation can be accessed on [public postman documenter](https://documenter.getpostman.com/view/19666540/2s93RNxEm4).
