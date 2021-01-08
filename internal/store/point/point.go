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
func (s *moneyStore) List(uid int64) (*core.UserPointDetail, error) {
	out := &core.UserPointDetail{}
	err := s.db.Select("*").
		Where("uid = ?", uid).
		Find(out).Error
	return out, err
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
