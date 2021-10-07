package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/uma-arai/sbcntr-backend/infrastructure"
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
			router.Logger.Fatal(router.Start(":80"))
		} else {
			router.Logger.Fatal(router.StartTLS(":443",
				os.Getenv(envTLSCert), os.Getenv(envTLSKey)))
		}
	}()

	<-quit
	fmt.Println("Caught SIGTERM, shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := router.Shutdown(ctx); err != nil {
		router.Logger.Fatal("Error:", err)
	}
	fmt.Println("Exited app")
}
