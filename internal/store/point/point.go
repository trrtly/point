package point

import (
	"point/internal/core"
	"point/internal/store/shared/db"

	"github.com/sirupsen/logrus"
)

// New returns a new UserStore.
func New(d *db.DB) core.UserPointDetailStore {
	return &moneyStore{d}
}

type moneyStore struct {
	db *db.DB
}

// List returns a user assets from the datastore.
func (s *moneyStore) List(
	r *core.UserPointDetailListRequest,
) ([]*core.UserPointDetail, int64, error) {
	var out []*core.UserPointDetail
	var count int64
	sdb := s.db.Model(&core.UserPointDetail{}).Where("uid = ?", r.UID)
	if r.Type > 0 {
		sdb.Where("type = ?", r.Type)
	}
	if r.FetchService {
		sdb.Where("service_point > ?", 0)
	} else {
		sdb.Where("money_point > ?", 0)
	}
	err := sdb.Scopes(db.Paginate(r.Page, r.PageSize)).Find(&out).Error
	if err != nil {
		return nil, 0, err
	}
	err = sdb.Count(&count).Error
	return out, count, err
}

// Create persists a new UserPointDetail in the db.
func (s *moneyStore) Create(m *core.UserPointDetail) error {
	res := s.db.Model(&core.UserPointDetail{}).Create(m)
	if res.Error != nil {
		logrus.WithFields(
			logrus.Fields{
				"data": m,
			},
		).Errorln("create UserPointDetail fail")
	}
	return res.Error
}

// BindUIDOpenid bind uid and openid.
func (s *moneyStore) BindUIDOpenid(uid, wechatUserID int64) error {
	upd := map[string]interface{}{
		"uid": uid,
	}
	return s.db.Debug().Model(&core.UserPointDetail{}).
		Where("wechat_user_id = ?", wechatUserID).Updates(upd).Error
}
