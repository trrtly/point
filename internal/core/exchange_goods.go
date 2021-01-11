package core

import "point/internal/core/trait"

type (
	// ExchangeGoods defines user_assets table
	ExchangeGoods struct {
		ID             uint32  `json:"id"`
		GoodsName      string  `json:"goods_name"`
		GoodsPic       string  `json:"goods_pic"`
		TotalNum       int32   `json:"total_num"`
		MoneyPoint     float64 `json:"money_point"`
		ServicePoint   float64 `json:"service_point"`
		GoodsType      float64 `json:"goods_type"`
		CreatedUserID  float64 `json:"created_user_id"`
		Desc           float64 `json:"desc"`
		UserDetailDesc float64 `json:"user_detail_desc"`
		CreatedTime    string  `json:"-"`
		ModifyTime     string  `json:"-"`
		trait.StartEndTime
	}

	// ExchangeGoodsStore defines operations for working with user_assets.
	ExchangeGoodsStore interface {
		// ListActivate returns a activate exchange goods list from the db
		ListActivate() ([]*ExchangeGoods, error)
		// Find returns a goods by id from the db
		Find(uint32) (*ExchangeGoods, error)
	}
)

// TableName defines the table name
func (ExchangeGoods) TableName() string {
	return "t_exchange_goods"
}
