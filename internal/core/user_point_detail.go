package core

import (
	"point/internal/core/trait"
)

type (
	// UserPointDetail defines user_point_detail table
	UserPointDetail struct {
		UID               int64  `json:"-"`
		Openid            string `json:"-"`
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
		// BindUIDOpenid bind uid and openid.
		BindUIDOpenid(int64, int64) error
	}
)

// TableName defines the user_point_detail table name in db
func (UserPointDetail) TableName() string {
	return "t_user_point_detail"
}
