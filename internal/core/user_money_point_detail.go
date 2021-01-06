package core

type (
	// UserMoneyPointDetail defines user_money_point_detail table
	UserMoneyPointDetail struct {
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

	// UserMoneyPointDetailStore defines operations for working with user_money_point_detail.
	UserMoneyPointDetailStore interface {
		// Find returns a user_money_point_detail from the db.
		List(int64) (*UserMoneyPointDetail, error)
		// Create persists a new UserMoneyPointDetail in the db.
		Create(*UserMoneyPointDetail) error
	}
	// UserMoneyPointDetailService provides access to user account
	UserMoneyPointDetailService interface {
		// Find returns the authenticated user.
		Find(access, refresh string) (*UserMoneyPointDetail, error)
	}
)

// TableName defines the user_money_point_detail table name in db
func (UserMoneyPointDetail) TableName() string {
	return "t_user_money_point_detail"
}
