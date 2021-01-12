package core

import (
	"point/internal/core/trait"
	"point/internal/pkg/hashids"

	"gorm.io/gorm"
)

type (
	// ExchangeGoods defines user_assets table
	ExchangeGoods struct {
		ID             int64   `json:"-"`
		GoodsYyid      string  `json:"goods_yyid"`
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

// AfterFind 查询后钩子
func (g *ExchangeGoods) AfterFind(tx *gorm.DB) (err error) {
	if g.ID > 0 {
		yyid, err := hashids.DefaultHd.EncodeInt64([]int64{g.ID})
		if err == nil {
			g.GoodsYyid = yyid
		}
	}
	return
}
