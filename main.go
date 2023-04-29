package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mustafa-533/rest-api/api"
	"github.com/mustafa-533/rest-api/db"
	"github.com/mustafa-533/rest-api/utils"
	"go.uber.org/zap"
)

type Config struct {
	LogLevel   string
	ListenHTTP string

	MySQLURI string
}

func main() {
	cfg := Config{
		LogLevel:   os.Getenv("LOG_LEVEL"),
		ListenHTTP: os.Getenv("LISTEN_HTTP_PORT"),

		MySQLURI: os.Getenv("MYSQL_URI"),
	}

	// Set Default Value for Port
	if cfg.ListenHTTP == "" {
		cfg.ListenHTTP = "8080"
	}

	logger, err := utils.LoadLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalln("Load Logger Error", err)
	}

	mysql, err := db.LoadMySqlDB(cfg.MySQLURI)
	if err != nil {
		logger.Fatal("Load MySQL Error")
	}

	// Cancel Context
	ctx, cancel := context.WithCancel(context.Background())

	var (
		handler = api.NewHandler(mysql, logger)

		// Load Routes
		router = handler.LoadRoutes()
	)

	server := &http.Server{
		Addr:         cfg.ListenHTTP,
		Handler:      router,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}

	logger.Info("Starting HTTP Server", zap.String("listen.http", cfg.ListenHTTP))

	go func() {
		// serve connections
		if srvErr := server.ListenAndServe(); srvErr != nil && !errors.Is(srvErr, http.ErrServerClosed) {
			logger.Fatal("Server Error", zap.Error(srvErr))
		}
	}()

	// Listen for Shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server Shutdown Error", zap.Error(err))
	}

	close(quit)
	cancel()
}
