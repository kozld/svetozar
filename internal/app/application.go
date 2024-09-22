package app

import (
	"context"

	"github.com/kozld/svetozar/internal/app/dependencies"
)

// Application is a main struct for the application that contains general information
type Application struct {
	Container *dependencies.Container
}

// InitializeApplication initializes new application
func InitializeApplication() (*Application, error) {
	return BuildApplication()
}

// Start starts application services
func (a *Application) Start(ctx context.Context, cli bool) {
	if cli {
		return
	}

	a.Container.Repository.MustMigrate()
}
