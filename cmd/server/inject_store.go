package main

import (
	"point/cmd/server/config"
	"point/internal/core"
	"point/internal/store/shared/db"
	"point/internal/store/assets"
	"point/internal/store/activity"

	"github.com/google/wire"
)

// wire set for loading the stores.
var storeSet = wire.NewSet(
	provideDatabase,
	provideAssetsStore,
	provideActivityStore,
)

// provideDatabase is a Wire provider function that provides a
// database connection, configured from the environment.
func provideDatabase(config config.Config) (*db.DB, error) {
	return db.Connect(&config.Database)
}

// provideAssetsStore is a Wire provider function that provides a
// user_assets datastore
func provideAssetsStore(db *db.DB) core.UserAssetsStore {
	assets := assets.New(db)
	return assets
}

// provideActivityStore is a Wire provider function that provides a
// activity datastore
func provideActivityStore(db *db.DB) core.ActivityStore {
	activity := activity.New(db)
	return activity
}
