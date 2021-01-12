package trait

import (
	"point/internal/pkg/hashids"

	"gorm.io/gorm"
)

type IDYyid struct {
	ID   int64  `json:"-"`
	Yyid string `json:"-"`
}

// AfterFind 查询后钩子
func (s *IDYyid) AfterFind(tx *gorm.DB) (err error) {
	if s.ID > 0 {
		yyid, err := hashids.DefaultHd.EncodeInt64([]int64{s.ID})
		if err == nil {
			s.Yyid = yyid
		}
	}
	return
}
