package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	apploader "github.com/auth-user-proxy-microservice/user-microservice/src/app-loader"
	"github.com/auth-user-proxy-microservice/user-microservice/src/config"
	"github.com/auth-user-proxy-microservice/user-microservice/src/model"
	"github.com/auth-user-proxy-microservice/user-microservice/src/services"
	"github.com/urfave/negroni"
)

func main() {

	configFile := flag.String("config", "config.json", "Configuration file in JSON-format")
	flag.Parse()

	if len(*configFile) > 0 {
		config.SetFilePath(*configFile)
	}

	// initializing global Application Service
	apploader.LoadApplicationServices()

	middlewareManager := negroni.New()
	middlewareManager.Use(negroni.NewRecovery())
	middlewareManager.UseHandler(services.NewRouter())

	server := &http.Server{
		Addr:         apploader.AppLoaderService.ConfigService.ListenURL,
		Handler:      middlewareManager,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	fmt.Printf("Starting application on port - %s", apploader.AppLoaderService.ConfigService.ListenURL)

	var err error
	if apploader.AppLoaderService.ConfigService.HTTPProtocol == model.HTTP {
		err = server.ListenAndServe()
	} else {
		err = server.ListenAndServeTLS(apploader.AppLoaderService.ConfigService.CertificateFile, apploader.AppLoaderService.ConfigService.CertificateKeyFile)
	}

	if err != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		e := server.Shutdown(ctx)
		log.Fatalf("Stop running application: %v, shutdown: %v", err, e)
	}
}
