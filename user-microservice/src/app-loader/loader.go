package apploader

import (
	"context"
	"log"
	"sync"

	"github.com/auth-user-proxy-microservice/user-microservice/src/config"
)

// AppLoader defines the whole configuration of the Application.
// It incorporates partial configurations of components
type AppLoader struct {
	ConfigService config.Configuration
}

// AppLoaderService is the whole configuration of the Application
var AppLoaderService *AppLoader

// LoadApplicationServices loads all partial configurations of components
// and populates the AppLoaderService with the configuration data
func LoadApplicationServices(ctx context.Context, wg *sync.WaitGroup) {
	loadFuncs := []func() error{
		config.Load,
	}

	for _, load := range loadFuncs {
		err := load()
		if err != nil {
			log.Fatal(err)
		}
	}

	AppLoaderService = &AppLoader{
		ConfigService: config.Config,
	}
}
