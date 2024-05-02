package internal

import (
	"go.uber.org/fx"
	"github.com/greeniet87/test_yokai/internal/repository"
)

// Register is used to register the application dependencies.
func Register() fx.Option {
	return fx.Options(
		fx.Provide(repository.NewExampleRepository),
	)
}
