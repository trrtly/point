package user

import (
	"point/internal/core"
)

type service struct {
	assets core.UserAssetsStore
}

// New returns a new User service that provides access to
// user data from the source code management system.
func New(
	assets core.UserAssetsStore,
) core.UserAssetsService {
	return &service{assets: assets}
}

func (s *service) Find(access, refresh string) (*core.UserAssets, error) {
	return nil, nil
}
