package main

import (
	"point/internal/service/user"

	"github.com/google/wire"
)

// wire set for loading the services.
var serviceSet = wire.NewSet(
	user.New,
)
