package core

import (
	"point/internal/core/trait"

	"gorm.io/gorm"
)

type (
	// UserPointDetail defines user_money_point_detail table
	UserPointDetail struct {
		trait.IDYyid
		UID               int64  `json:"-"`
		ActivityID        int32  `json:"-"`
		ActivitySpecialID int32  `json:"-"`
		GoodsID           int64  `json:"-"`
		GoodsNum          int32  `json:"-"`
		Type              int8   `json:"type"`
		Status            int8   `json:"status"`
		Desc              string `json:"desc"`
		trait.CreatedUpdatedTime
		trait.Point
	}

	UserPointDetailListRequest struct {
		UID          int64 `query:"uid,required,number"`
		FetchService bool  `query:"fetchService,number"`
		Page         int   `query:"page,number"`
		PageSize     int   `query:"pageSize,number"`
		Type         int8  `query:"type,number"`
	}

	// UserPointDetailStore defines operations for working with user_money_point_detail.
	UserPointDetailStore interface {
		// Find returns a user_money_point_detail from the db.
		List(*UserPointDetailListRequest) ([]*UserPointDetail, int64, error)
		// Create persists a new UserPointDetail in the db.
		Create(*UserPointDetail) error
	}
)

// TableName defines the user_money_point_detail table name in db
func (UserPointDetail) TableName() string {
	return "t_user_point_detail"
}

// AfterFind 查询后钩子
func (s *UserPointDetail) AfterFind(tx *gorm.DB) (err error) {
	err = s.IDYyid.AfterFind(tx)
	s.FormatTime()
	return
}
