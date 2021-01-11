package goods

import (
	"point/internal/core"
	"point/internal/core/status"
	"point/internal/store/shared/db"
	"time"
)

// New returns a new ExchangeGoodsStore.
func New(db *db.DB) core.ExchangeGoodsStore {
	return &goodsStore{db}
}

type goodsStore struct {
	db *db.DB
}

// ListActivate returns a activate exchange goods list from the db.
func (s *goodsStore) ListActivate() ([]*core.ExchangeGoods, error) {
	var out []*core.ExchangeGoods
	now := time.Now().Format("2006-01-02 15:04:05")
	err := s.db.Select("*").
		Where("start_time >= ?", now).
		Where("end_time < ?", now).
		Where("status = ?", status.Regular).
		Find(out).Error
	return out, err
}

// Find returns a goods by id from the db.
func (s *goodsStore) Find(id uint32) (*core.ExchangeGoods, error) {
	out := &core.ExchangeGoods{}
	err := s.db.Select("*").
		Where("id = ?", id).
		First(out).Error
	return out, err
}
