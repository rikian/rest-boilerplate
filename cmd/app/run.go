package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang-test/cmd/migrations"
	"golang-test/config"
	h "golang-test/src/handler"
	"golang-test/src/helper"
	"golang-test/src/repositories"
	"golang-test/src/service/a"
	"golang-test/src/service/e"
	"golang-test/src/service/r"

	"github.com/go-playground/validator/v10"
)

func Run() {
	// load env file
	config.LoadEnvFile()
	// build zap logger
	logger := config.BuildLogger()
	// create databse connection
	db := config.ConnectDB()
	// initial service

	jwt := helper.NewJwtImpl([]byte(os.Getenv("JWT_SECRET")))
	validate := helper.NewValidation(validator.New())
	repo := repositories.NewRepo(logger)
	serviceR := r.NewResponseHandler(logger)
	serviceE := e.NewErrorHandler(serviceR, logger)
	serviceA := a.NewAService(logger, repo, db, validate, serviceR)
	handler := h.NewHandler(logger, serviceE, serviceA, jwt)
	// create server
	server := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: handler,
	}

	// run migration
	migrations.RunMigration(db)

	claim, e := jwt.GenerateJWT("rikianfaisal", 3600000)
	if e != nil {
		log.Fatal("failed generate jwt")
	}

	log.Print("your jwt for auth : " + claim)

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		log.Print("Server running at " + os.Getenv("PORT"))
		logger.Info("Server running at " + os.Getenv("PORT"))
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("Server error:", err)
		}
	}()

	// Block until a signal is received
	<-done
	log.Print("Shutting down server gracefully...")
	logger.Info("Shutting down server gracefully...")

	// Create a context with a timeout to allow for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Stop accepting new connections
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server shutdown error:", err)
	}

	log.Print("Server gracefully stopped.")
	logger.Info("Server gracefully stopped.")
}
