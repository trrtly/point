package point

import (
	"point/internal/core"
	"point/internal/store/shared/db"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
	sdb.Order("id desc")
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
		).Errorln("create UserPointDetail fail", res.Error)
	}
	return res.Error
}

// BindUIDWechatUID bind uid and wechat user id.
func (s *moneyStore) BindUIDWechatUID(uid, wechatUserID int64) error {
	upd := map[string]interface{}{
		"uid": uid,
	}
	return s.db.Model(&core.UserPointDetail{}).
		Where("wechat_user_id = ?", wechatUserID).Updates(upd).Error
}

// HasBindUIDWechatUID find a record by uid and wechat user id.
func (s *moneyStore) HasBindUIDWechatUID(wechatUserID int64) bool {
	out := &core.UserPointDetail{}
	err := s.db.Model(&core.UserPointDetail{}).
		Where("wechat_user_id = ?", wechatUserID).
		Where("uid > ?", 0).
		First(out).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

// FindMoneyServicePointSum find a record by uid and wechat user id.
func (s *moneyStore) FindMoneyServicePointSum(wechatUserID int64) (float64, float64, error) {
	type result struct {
		Money   float64
		Service float64
	}
	out := &result{}
	err := s.db.Model(&core.UserPointDetail{}).
		Select("sum(service_point) as Service, sum(money_point) as Money").
		Where("wechat_user_id = ?", wechatUserID).
		First(out).Error
	return out.Money, out.Service, err
}
