package trait

import (
	"time"
	"point/internal/core/status"
)

type (
	// StartEndTime defines a StartEndTime is alife
	StartEndTime struct {
		StartTime time.Time `json:"-"`
		EndTime   time.Time `json:"-"`
		Status    int8      `json:"-"`
	}
)

// IsRegular defines the StartEndTime record is regular
func (a *StartEndTime) IsRegular() bool {
	return a.Status == status.Regular
}

// IsStarted defines the StartEndTime is started
func (a *StartEndTime) IsStarted() bool {
	today := time.Now()
	return today.After(a.StartTime)
}

// IsNotEnded defines the StartEndTime is not ended
func (a *StartEndTime) IsNotEnded() bool {
	if a.EndTime.IsZero() {
		return true
	}
	today := time.Now()
	return today.Before(a.EndTime)
}

// IsActivite defines the StartEndTime is activite
func (a *StartEndTime) IsActivite() bool {
	return a.IsRegular() && a.IsStarted() && a.IsNotEnded()
}
