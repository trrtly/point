package trait

import (
	"point/internal/pkg/hd"

	"gorm.io/gorm"
)

type IDYyid struct {
	ID   int64  `json:"-"`
	Yyid string `json:"yyid" gorm:"-"`
}

// AfterFind 查询后钩子
func (s *IDYyid) AfterFind(tx *gorm.DB) (err error) {
	if s.ID > 0 {
		yyid, err := hd.DefaultHd.EncodeInt64([]int64{s.ID})
		if err == nil {
			s.Yyid = yyid
		}
	}
	return
}
