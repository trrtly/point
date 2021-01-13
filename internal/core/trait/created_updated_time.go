package trait

import (
	// "fmt"
	"time"
)

type (
	// CreatedUpdatedTime defines a StartEndTime is alife
	CreatedUpdatedTime struct {
		CreatedAt       time.Time `json:"-"`
		UpdatedAt       time.Time `json:"-"`
		CreatedAtString string    `json:"created_at"`
		UpdatedAtString string    `json:"updated_at"`
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
