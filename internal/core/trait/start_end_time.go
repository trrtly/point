package trait

import (
	"time"
)

type (
	// StartEndTime defines a StartEndTime is alife
	StartEndTime struct {
		StartTime time.Time `json:"-"`
		EndTime   time.Time `json:"-"`
		Status
	}
)

// IsStarted defines the StartEndTime is started
func (a *StartEndTime) IsStarted() bool {
	return time.Now().After(a.StartTime)
}

// IsNotEnded defines the StartEndTime is not ended
func (a *StartEndTime) IsNotEnded() bool {
	if a.EndTime.IsZero() {
		return true
	}
	return time.Now().Before(a.EndTime)
}

// IsActivite defines the StartEndTime is activite
func (a *StartEndTime) IsActivite() bool {
	return a.IsRegular() && a.IsStarted() && a.IsNotEnded()
}
