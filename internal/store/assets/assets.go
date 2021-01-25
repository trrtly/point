package assets

import (
	"point/internal/core"
	"point/internal/store/shared/db"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// New returns a new UserStore.
func New(d *db.DB) core.UserAssetsStore {
	return &assetsStore{d}
}

type assetsStore struct {
	db *db.DB
}

// Find returns a user assets from the datastore.
func (s *assetsStore) Find(uid int64) (*core.UserAssets, error) {
	out := &core.UserAssets{}
	err := s.db.Model(&core.UserAssets{}).
		Where("uid = ?", uid).
		First(out).Error
	return out, err
}

// IncrPoint increment a user's money or service point or both
func (s *assetsStore) IncrPoint(uid int64, moneyPoint float64, servicePoint float64) error {
	if uid <= 0 {
		return nil
	}
	upd := map[string]interface{}{
		"money_point":   gorm.Expr("money_point + ?", moneyPoint),
		"service_point": gorm.Expr("service_point + ?", servicePoint),
	}
	return s.db.Model(&core.UserAssets{}).Where("uid = ?", uid).Updates(upd).Error
}

// DecrPoint decrement a user's money or service point or both
func (s *assetsStore) DecrPoint(uid int64, old, new *core.UserAssets) error {
	upd := map[string]interface{}{
		"money_point":   gorm.Expr("money_point - ?", new.MoneyPoint),
		"service_point": gorm.Expr("service_point - ?", new.ServicePoint),
	}
	return s.db.Model(&core.UserAssets{}).
		Where("uid = ?", uid).
		Where("money_point = ?", old.MoneyPoint).
		Where("service_point = ?", old.ServicePoint).
		Updates(upd).Error
}

// Create persists a new UserAssets in the db.
func (s *assetsStore) Create(m *core.UserAssets) error {
	res := s.db.Model(&core.UserAssets{}).Create(m)
	if res.Error != nil {
		logrus.WithFields(
			logrus.Fields{
				"data": m,
			},
		).Errorln("create UserAssets fail", res.Error)
	}
	return res.Error
}
