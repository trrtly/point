package core

import (
	"point/internal/core/trait"
)

type (
	// ExchangeGoods defines user_assets table
	ExchangeGoods struct {
		trait.IDYyid
		// 商品名称
		GoodsName string `json:"goods_name" example:"兑换现金100"`
		// 商品图片
		GoodsPic string `json:"goods_pic" example:"http://coupons.quanduogo.com/ico.png"`
		TotalNum int32  `json:"-"`
		// 消费积分
		MoneyPoint float64 `json:"money_point" example:"100.00"`
		// 服务积分
		ServicePoint float64 `json:"service_point" example:"100.00"`
		// 商品类型 1表示 现金 2表示实物  3表示虚拟
		GoodsType     int8   `json:"goods_type" example:"1"`
		CreatedUserID string `json:"-"`
		// 商品描述
		Desc           string `json:"desc" example:"每100的服务积分和100个消费积分可兑换100块钱"`
		UserDetailDesc string `json:"-"`
		CreatedTime    string `json:"-"`
		ModifyTime     string `json:"-"`
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
