package core

import "point/internal/core/trait"

type (
	// ExchangeGoodsOrder defines user_assets table
	ExchangeGoodsOrder struct {
		ID        uint32 `json:"id"`
		UID       int64  `json:"uid"`
		GoodsID   int64  `json:"goods_id"`
		GoodsNum  uint32 `json:"goods_num"`
		Status    int8   `json:"status"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		trait.Point
		trait.StartEndTime
	}

	// ExchangeGoodsOrderStore defines operations for working with user_assets.
	ExchangeGoodsOrderStore interface {
		// Create persists a exchange goods order to the db.
		Create(*ExchangeGoodsOrder) error
	}
)

// TableName defines the table name
func (ExchangeGoodsOrder) TableName() string {
	return "t_exchange_goods_order"
}
