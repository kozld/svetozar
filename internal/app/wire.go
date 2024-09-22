//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/kozld/svetozar/config"
	"github.com/kozld/svetozar/internal/app/dependencies"
	"github.com/kozld/svetozar/internal/app/initializers"
	repository "github.com/kozld/svetozar/internal/repositories/message"
	"github.com/kozld/svetozar/internal/services/listener"
	"github.com/kozld/svetozar/internal/services/validator"
)

func BuildApplication() (*Application, error) {
	wire.Build(
		config.LoadConfig,
		repository.New,
		listener.New,
		validator.New,

		initializers.InitializeLogs,
		initializers.InitializeWorkerpool,
		initializers.InitializeDatabase,
		initializers.InitializeTDLibClient,

		wire.Struct(new(dependencies.Container), "*"),
		wire.Struct(new(Application), "*"),
	)

	return &Application{}, nil
}
