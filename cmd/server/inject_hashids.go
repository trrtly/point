package main

import (
	"point/cmd/server/config"
	"point/internal/pkg/hd"

	"github.com/google/wire"
)

// wire set for generate hashids instance.
var hashidsSet = wire.NewSet(
	provideHashids,
)

// provideDatabase is a Wire provider function that provides a
// database connection, configured from the environment.
func provideHashids(config config.Config) (*hd.HD, error) {
	hashid, err := hd.New(&config.Hashids)
	hd.DefaultHd = hashid
	return hashid, err
}
