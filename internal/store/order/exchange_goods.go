package order

import (
	"point/internal/core"
	"point/internal/store/shared/db"

	"github.com/sirupsen/logrus"
)

// New returns a new ExchangeexchangeGoodsStore.
func New(db *db.DB) core.ExchangeGoodsOrderStore {
	return &exchangeGoodsStore{db}
}

type exchangeGoodsStore struct {
	db *db.DB
}

// Create returns a activate exchange goods list from the db.
func (s *exchangeGoodsStore) Create(m *core.ExchangeGoodsOrder) error {
	res := s.db.Create(m)
	if res.Error != nil {
		logrus.WithFields(
			logrus.Fields{
				"id":            m.ID,
				"uid":           m.UID,
				"goods_id":      m.GoodsID,
				"goods_num":     m.GoodsNum,
				"money_point":   m.MoneyPoint,
				"service_point": m.ServicePoint,
				"status":        m.Status,
				"created_at":    m.CreatedAt,
				"updated_at":    m.UpdatedAt,
			},
		).Errorln("create exchange_goods_order fail")
	}
	return res.Error
}
