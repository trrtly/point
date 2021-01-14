package trait

import (
	"gorm.io/gorm"
)

type (
	// IDYyidCreatedUpdatedTime defines a StartEndTime is alife
	IDYyidCreatedUpdatedTime struct {
		IDYyid
		CreatedUpdatedTime
	}
)

// AfterFind 查询后钩子
func (s *IDYyidCreatedUpdatedTime) AfterFind(tx *gorm.DB) (err error) {
	s.IDYyid.AfterFind(tx)
	s.CreatedUpdatedTime.FormatTime()
	return
}
