package point

import (
	"point/internal/core"
	"point/internal/store/shared/db"

	"github.com/sirupsen/logrus"
)

// New returns a new UserStore.
func New(db *db.DB) core.UserPointDetailStore {
	return &moneyStore{db}
}

type moneyStore struct {
	db *db.DB
}

// List returns a user assets from the datastore.
func (s *moneyStore) List(uid int64) ([]*core.UserPointDetail, int64, error) {
	var out []*core.UserPointDetail
	var count int64
	sdb := s.db.Select("*").Where("uid = ?", uid)
	err := sdb.Find(&out).Error
	if err != nil {
		return nil, 0, err
	}
	err = sdb.Count(&count).Error
	return out, 0, err
}

// Create persists a new UserPointDetail in the db.
func (s *moneyStore) Create(m *core.UserPointDetail) error {
	res := s.db.Create(m)
	if res.Error != nil {
		logrus.WithFields(
			logrus.Fields{
				"id":            m.ID,
				"uid":           m.UID,
				"activity_id":   m.ActivityID,
				"goods_id":      m.GoodsID,
				"goods_num":     m.GoodsNum,
				"service_point": m.ServicePoint,
				"money_point":   m.MoneyPoint,
				"type":          m.Type,
				"status":        m.Status,
				"desc":          m.Desc,
				"created_at":    m.CreatedAt,
			},
		).Errorln("create UserPointDetail fail")
	}
	return res.Error
}
