//+build wireinject

package main

import (
	"point/cmd/server/config"
	"github.com/google/wire"
)

func InitializeApplication(config config.Config) (application, error) {
	wire.Build(
		serverSet,
		storeSet,
		newApplication,
	)
	return application{}, nil
}
