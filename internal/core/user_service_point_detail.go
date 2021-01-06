package core

type (
	// UserServicePointDetail defines user_money_point_detail table
	UserServicePointDetail struct {
		ID         int32   `json:"-"`
		Yyid       string  `json:"yyid"`
		Openid     string  `json:"-"`
		UID        int64   `json:"-"`
		ActivityID int32   `json:"-"`
		GoodsID    int32   `json:"-"`
		GoodsNum   int32   `json:"-"`
		Point      float64 `json:"point"`
		Type       int8    `json:"type"`
		Status     int8    `json:"status"`
		Desc       string  `json:"desc"`
		CreatedAt  string  `json:"created_at"`
	}

	// UserServicePointDetailStore defines operations for working with user_money_point_detail.
	UserServicePointDetailStore interface {
		// Find returns a user_money_point_detail from the db.
		List(int64) (*UserServicePointDetail, error)
		// Create persists a new UserServicePointDetail in the db.
		Create(*UserServicePointDetail) error
	}
	// UserServicePointDetailService provides access to user account
	UserServicePointDetailService interface {
		// Find returns the authenticated user.
		Find(access, refresh string) (*UserServicePointDetail, error)
	}
)

// TableName defines the user_money_point_detail table name in db
func (UserServicePointDetail) TableName() string {
	return "t_user_service_point_detail"
}
