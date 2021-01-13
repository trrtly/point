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
func (s *moneyStore) List(uid int64, ptype int8, page, pageSize int) ([]*core.UserPointDetail, int64, error) {
	var out []*core.UserPointDetail
	var count int64
	sdb := s.db.Where("uid = ?", uid).Where("point_type", ptype)
	err := sdb.Scopes(db.Paginate(page, pageSize)).Find(&out).Error
	if err != nil {
		return nil, 0, err
	}
	err = sdb.Count(&count).Error
	return out, count, err
}

// Create persists a new UserPointDetail in the db.
func (s *moneyStore) Create(m *core.UserPointDetail) error {
	res := s.db.Create(m)
	if res.Error != nil {
		logrus.WithFields(
			logrus.Fields{
				"data": m,
			},
		).Errorln("create UserPointDetail fail")
	}
	return res.Error
}
