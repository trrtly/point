package core

import (
	"point/internal/core/trait"
)

type (
	// UserPointDetail defines user_point_detail table
	UserPointDetail struct {
		UID               int64  `json:"-"`
		WechatUserID      int64  `json:"-"`
		ActivityID        int32  `json:"-"`
		ActivitySpecialID int32  `json:"-"`
		GoodsID           int64  `json:"-"`
		GoodsNum          int32  `json:"-"`
		Type              int8   `json:"type"`
		Status            int8   `json:"status"`
		Desc              string `json:"desc"`
		trait.IDYyidCreatedUpdatedTime
		trait.MoneyServicePoint
	}

	// UserPointDetailListRequest defines api/point/details request data struct
	UserPointDetailListRequest struct {
		UID          int64 `query:"uid,required,number"`
		FetchService bool  `query:"fetchService,number"`
		Page         int   `query:"page,number"`
		PageSize     int   `query:"pageSize,number"`
		Type         int8  `query:"type,number"`
	}

	// UserPointDetailStore defines operations for working with user_point_detail.
	UserPointDetailStore interface {
		// Find returns a user_point_detail from the db.
		List(*UserPointDetailListRequest) ([]*UserPointDetail, int64, error)
		// Create persists a new user_point_detail record in the db.
		Create(*UserPointDetail) error
		// BindUIDWechatUID bind uid and wechat_user_id.
		BindUIDWechatUID(int64, int64) error
		// HasBindUIDWechatUID find a record by uid and wechat_user_id.
		HasBindUIDWechatUID(int64) bool
		// FindMoneyServicePointSum find money service point sum by wechat_user_id.
		FindMoneyServicePointSum(int64) (float64, float64, error)
		// CountActivityNumPreDay .
		CountActivityNumPreDay(uid int64, activity *Activity) (int64, error)
		// CountSpecialNumPreDay .
		CountSpecialNumPreDay(uid int64, special *ActivitySpecial) (int64, error)
	}
)

const (
	// UserPointDetailDeleted 已删除
	UserPointDetailDeleted = 0
	// UserPointDetailApply 申请中
	UserPointDetailApply = 1
	// UserPointDetailArrived 已到账
	UserPointDetailArrived = 2
	// UserPointDetailExchanged 已兑换
	UserPointDetailExchanged = 3
)

// TableName defines the user_point_detail table name in db
func (UserPointDetail) TableName() string {
	return "t_user_point_detail"
}
