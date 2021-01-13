package goods

import (
	"point/internal/core"
	"point/internal/core/status"
	"point/internal/store/shared/db"
	"time"
)

// New returns a new ExchangeGoodsStore.
func New(d *db.DB) core.ExchangeGoodsStore {
	m := d.Model(&core.ExchangeGoods{})
	return &goodsStore{&db.DB{DB: m}}
}

type goodsStore struct {
	db *db.DB
}

// ListActivate returns a activate exchange goods list from the db.
func (s *goodsStore) ListActivate(page, pageSize int) ([]*core.ExchangeGoods, int64, error) {
	var out []*core.ExchangeGoods
	var count int64
	now := time.Now().Format("2006-01-02 15:04:05")
	sdb := s.db.Where("start_time <= ?", now).
		Where("end_time > ?", now).
		Where("status = ?", status.Regular)

	err := sdb.Scopes(db.Paginate(page, pageSize)).Find(&out).Error
	if err != nil {
		return nil, 0, err
	}
	err = sdb.Count(&count).Error

	return out, count, err
}

// Find returns a goods by id from the db.
func (s *goodsStore) Find(id uint32) (*core.ExchangeGoods, error) {
	out := &core.ExchangeGoods{}
	err := s.db.Select("*").
		Where("id = ?", id).
		First(out).Error
	return out, err
}
