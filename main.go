package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sferawann/test_mnc/config"
	"github.com/sferawann/test_mnc/controller"
	"github.com/sferawann/test_mnc/repository"
	"github.com/sferawann/test_mnc/router"
	"github.com/sferawann/test_mnc/usecase"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	//init repository
	userRepo := repository.NewUserRepoImpl("json/user.json")
	accRepo := repository.NewAccountRepoImpl("json/account.json")
	hisRepo := repository.NewHistoryRepoImpl("json/history.json")
	traRepo := repository.NewTransferRepoImpl("json/transfer.json")
	sesRepo := repository.NewSessionRepoImpl("json/session.json")

	//init usecase
	userUsecase := usecase.NewUserUsecaseImpl(userRepo)
	accUsecase := usecase.NewAccountUsecaseImpl(accRepo, userRepo)
	hisUsecase := usecase.NewHistoryUsecaseImpl(hisRepo, userRepo, accRepo)
	traUsecase := usecase.NewTransferUsecaseImpl(traRepo, userRepo, accRepo, hisRepo)
	sesUsecase := usecase.NewSessionUsecaseImpl(sesRepo, userRepo)
	authUsecase := usecase.NewAuthUsecaseImpl(userRepo, sesRepo)

	//init controller
	userCon := controller.NewUserController(userUsecase)
	accCon := controller.NewAccountController(accUsecase)
	hisCon := controller.NewHistoryController(hisUsecase)
	traCon := controller.NewTransferController(traUsecase, accUsecase)
	sesCon := controller.NewSessionController(sesUsecase)
	authCon := controller.NewAuthController(authUsecase)

	//init routes
	routes := router.NewRouter(userCon, accCon, hisCon, traCon, sesCon, authCon)
	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if server_err := server.ListenAndServe(); err != nil {
		log.Fatal(server_err)
	}
}
