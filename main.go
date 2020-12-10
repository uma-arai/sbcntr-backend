package main

import (
	"github.com/iselegant/cnappdemo/infrastructure"
	"os"
)

const (
	envTLSCert = "TLS_CERT"
	envTLSKey  = "TLS_KEY"
)

func main() {
	router := infrastructure.Router()

	if os.Getenv(envTLSCert) == "" || os.Getenv(envTLSKey) == "" {
		router.Logger.Fatal(router.Start(":80"))
	} else {
		router.Logger.Fatal(router.StartTLS(":443",
			os.Getenv(envTLSCert), os.Getenv(envTLSKey)))
	}
}
