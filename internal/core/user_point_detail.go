package core

import (
	"point/internal/core/trait"
)

type (
	// UserPointDetail defines user_money_point_detail table
	UserPointDetail struct {
		trait.IDYyid
		Yyid              string `json:"detail_yyid"`
		UID               int64  `json:"-"`
		ActivityID        int32  `json:"-"`
		ActivitySpecialID int32  `json:"-"`
		GoodsID           int32  `json:"-"`
		GoodsNum          int32  `json:"-"`
		Type              int8   `json:"type"`
		Status            int8   `json:"status"`
		Desc              string `json:"desc"`
		CreatedAt         string `json:"created_at"`
		trait.Point
	}

	// UserPointDetailStore defines operations for working with user_money_point_detail.
	UserPointDetailStore interface {
		// Find returns a user_money_point_detail from the db.
		List(int64) ([]*UserPointDetail, int64, error)
		// Create persists a new UserPointDetail in the db.
		Create(*UserPointDetail) error
	}
)

// TableName defines the user_money_point_detail table name in db
func (UserPointDetail) TableName() string {
	return "t_user_point_detail"
}
