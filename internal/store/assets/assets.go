package assets

import (
	"point/internal/core"
	"point/internal/store/shared/db"
)

// New returns a new UserStore.
func New(db *db.DB) core.UserAssetsStore {
	return &assetsStore{db}
}

type assetsStore struct {
	db *db.DB
}

// Find returns a user assets from the datastore.
func (s *assetsStore) Find(uid int64) (*core.UserAssets, error) {
	out := &core.UserAssets{}
	err := s.db.Select("id,uid,money_point,service_point").
		Where("uid = ?", uid).
		First(&out).Error
	return out, err
}
