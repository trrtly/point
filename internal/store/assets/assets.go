package assets

import (
	"point/internal/core"
	"point/internal/store/shared/db"

	"gorm.io/gorm"
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
		First(out).Error
	return out, err
}

// IncrPoint increment a user's money or service point or both
func (s *assetsStore) IncrPoint(uid int64, moneyPoint float64, servicePoint float64) error {
	upd := map[string]interface{}{
		"money_point":   gorm.Expr("money_point + ?", moneyPoint),
		"service_point": gorm.Expr("service_point + ?", servicePoint),
	}
	return s.db.Where("uid = ?", uid).Updates(upd).Error
}

// DecrPoint decrement a user's money or service point or both
func (s *assetsStore) DecrPoint(uid int64, moneyPoint float64, servicePoint float64) error {
	upd := map[string]interface{}{
		"money_point":   gorm.Expr("money_point - ?", moneyPoint),
		"service_point": gorm.Expr("service_point - ?", servicePoint),
	}
	return s.db.Where("uid = ?", uid).Updates(upd).Error
}
