package main

import (
	"point/cmd/server/config"
	"point/internal/pkg/hashids"

	"github.com/google/wire"
)

// wire set for generate hashids instance.
var hashidsSet = wire.NewSet(
	provideHashids,
)

// provideDatabase is a Wire provider function that provides a
// database connection, configured from the environment.
func provideHashids(config config.Config) (*hashids.HD, error) {
	hd, err := hashids.New(&config.Hashids)
	hashids.DefaultHd = hd
	return hd, err
}
