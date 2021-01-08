package main

import (
	"point/cmd/server/config"
	"point/internal/core"
	"point/internal/store/activity"
	"point/internal/store/assets"
	"point/internal/store/point"
	"point/internal/store/shared/db"
	"point/internal/store/special"

	"github.com/google/wire"
)

// wire set for loading the stores.
var storeSet = wire.NewSet(
	provideDatabase,
	provideAssetsStore,
	provideActivityStore,
	provideActivitySpecialStore,
	providePointDetailStore,
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

// provideActivitySpecialStore is a Wire provider function that provides a
// activity_special datastore
func provideActivitySpecialStore(db *db.DB) core.ActivitySpecialStore {
	special := special.New(db)
	return special
}

// providePointDetailStore is a Wire provider function that provides a
// user_point_detail datastore
func providePointDetailStore(db *db.DB) core.UserPointDetailStore {
	detail := point.New(db)
	return detail
}
