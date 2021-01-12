package core

import (
	"point/internal/core/trait"
)

type (
	// ExchangeGoods defines user_assets table
	ExchangeGoods struct {
		trait.IDYyid
		GoodsName      string  `json:"goods_name"`
		GoodsPic       string  `json:"goods_pic"`
		TotalNum       int32   `json:"-"`
		MoneyPoint     float64 `json:"money_point"`
		ServicePoint   float64 `json:"service_point"`
		GoodsType      int8    `json:"goods_type"`
		CreatedUserID  string  `json:"-"`
		Desc           string  `json:"desc"`
		UserDetailDesc string  `json:"-"`
		CreatedTime    string  `json:"-"`
		ModifyTime     string  `json:"-"`
		trait.StartEndTime
	}

	// ExchangeGoodsStore defines operations for working with user_assets.
	ExchangeGoodsStore interface {
		// ListActivate returns a activate exchange goods list from the db
		ListActivate(int, int) ([]*ExchangeGoods, int64, error)
		// Find returns a goods by id from the db
		Find(uint32) (*ExchangeGoods, error)
	}
)

// TableName defines the table name
func (ExchangeGoods) TableName() string {
	return "t_exchange_goods"
}
