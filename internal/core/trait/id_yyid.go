package trait

import (
	"point/internal/pkg/hd"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IDYyid struct {
	ID   int64  `json:"-"`
	Yyid string `json:"yyid,omitempty" gorm:"-"`
}

// AfterFind 查询后钩子
func (s *IDYyid) AfterFind(tx *gorm.DB) (err error) {
	if s.ID > 0 {
		yyid, err := hd.DefaultHd.EncodeInt64([]int64{s.ID})
		s.Yyid = yyid
		if err != nil {
			logrus.WithFields(
				logrus.Fields{
					"data": s,
				},
			).Errorln("default hashid 生成失败", err)
		}
	}
	return
}
