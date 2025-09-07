package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/horsewin/echo-playground-v2/infrastructure"
	"github.com/rs/zerolog/log"
)

const (
	envTLSCert = "TLS_CERT"
	envTLSKey  = "TLS_KEY"
)

func main() {
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)

	router := infrastructure.Router()

	// Start server
	go func() {
		if os.Getenv(envTLSCert) == "" || os.Getenv(envTLSKey) == "" {
			log.Info().Msg("Starting server on :8081")
			if err := router.Start(":8081"); err != nil {
				log.Fatal().Err(err).Msg("Failed to start server")
			}
		} else {
			log.Info().Msg("Starting server with TLS on :443")
			if err := router.StartTLS(":443",
				os.Getenv(envTLSCert), os.Getenv(envTLSKey)); err != nil {
				log.Fatal().Err(err).Msg("Failed to start server with TLS")
			}
		}
	}()

	<-quit
	log.Info().Msg("Caught SIGTERM, shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := router.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Error during server shutdown")
	}
	log.Info().Msg("Exited app")
}
