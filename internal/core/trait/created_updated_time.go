package trait

import (
	// "fmt"
	"time"
)

type (
	// CreatedUpdatedTime defines a StartEndTime is alife
	CreatedUpdatedTime struct {
		CreatedAt       time.Time `json:"-" gorm:"autoCreateTime"`
		UpdatedAt       time.Time `json:"-" gorm:"autoUpdateTime"`
		CreatedAtString string    `json:"created_at" gorm:"-"`
		UpdatedAtString string    `json:"updated_at" gorm:"-"`
	}
)

func (s *CreatedUpdatedTime) FormatTime() {
	s.CreatedAtString = s.format(s.CreatedAt)
	s.UpdatedAtString = s.format(s.UpdatedAt)
}

func (s *CreatedUpdatedTime) format(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return time.Time(t).Format("2006-01-02 15:04:05")
}
