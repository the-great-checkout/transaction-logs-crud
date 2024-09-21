package main

import (
	"github.com/Netflix/go-env"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/the-great-checkout/transaction-logs-crud/internal/controller"
	"github.com/the-great-checkout/transaction-logs-crud/internal/mapper"
	"github.com/the-great-checkout/transaction-logs-crud/internal/repository"
	"github.com/the-great-checkout/transaction-logs-crud/internal/service"
	_ "github.com/the-great-checkout/transactions-crud/docs"

	"github.com/the-great-checkout/transaction-logs-crud/internal/database"
)

type Environment struct {
	Postgres struct {
		Schema string `env:"POSTGRES_SCHEMA,default=transaction_logs."`
		DSN    string `env:"POSTGRES_DSN,default=host=localhost user=user password=password dbname=postgres port=5432 sslmode=disable"`
	}

	Port string `env:"PORT,default=:8082"`
}

// @title Transaction Logs CRUD API
// @version 1.0
// @description This is a sample server for transaction log CRUD.
// @host localhost:8082
// @BasePath /

func main() {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		panic(err)
	}

	postgres := database.NewPostgres(environment.Postgres.DSN, environment.Postgres.Schema)

	transactionRepository := repository.NewTransactionLogRepository(postgres)
	transactionMapper := mapper.NewTransactionLogMapper()
	transactionService := service.NewTransactionLogService(transactionRepository, transactionMapper)
	transactionController := controller.NewTransactionLogController(transactionService)

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/v1")

	// Route for creating a new transaction log
	v1.POST("/transactions/:transactionId/logs", transactionController.CreateHandler)
	// Route for retrieving all logs by transaction ID
	v1.GET("/transactions/:transactionId/logs", transactionController.FindAllByTransactionIDHandler)

	e.Logger.Fatal(e.Start(environment.Port))
}
