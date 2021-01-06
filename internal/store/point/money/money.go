package money

import (
	"point/internal/core"
	"point/internal/store/shared/db"

	"github.com/sirupsen/logrus"
)

// New returns a new UserStore.
func New(db *db.DB) core.UserMoneyPointDetailStore {
	return &moneyStore{db}
}

type moneyStore struct {
	db *db.DB
}

// List returns a user assets from the datastore.
func (s *moneyStore) List(uid int64) (*core.UserMoneyPointDetail, error) {
	out := &core.UserMoneyPointDetail{}
	err := s.db.Select("*").
		Where("uid = ?", uid).
		Find(&out).Error
	return out, err
}

// Create persists a new UserMoneyPointDetail in the db.
func (s *moneyStore) Create(m *core.UserMoneyPointDetail) error {
	res := s.db.Create(m)
	if res.Error != nil {
		logrus.WithFields(
			logrus.Fields{
				"id":          m.ID,
				"yyid":        m.Yyid,
				"openid":      m.Openid,
				"uid":         m.UID,
				"activity_id": m.ActivityID,
				"goods_id":    m.GoodsID,
				"goods_num":   m.GoodsNum,
				"point":       m.Point,
				"type":        m.Type,
				"status":      m.Status,
				"desc":        m.Desc,
				"created_at":  m.CreatedAt,
			},
		).Errorln("create UserMoneyPointDetail fail")
	}
	return res.Error
}
