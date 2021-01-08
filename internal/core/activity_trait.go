package core

import (
	"time"
)

type (
	// ActivityTrait defines a activity is alife
	ActivityTrait struct {
		PointTrait
		StartTime time.Time `json:"start_time"`
		EndTime   time.Time `json:"end_time"`
		Status    int8      `json:"status"`
		NumPreDay int32     `json:"num_pre_day"`
		NumTotal  int32     `json:"num_total"`
	}
)

// IsRegular defines the activity record is regular
func (a *ActivityTrait) IsRegular() bool {
	return a.Status == StatusRegular
}

// IsStarted defines the activity is started
func (a *ActivityTrait) IsStarted() bool {
	today := time.Now()
	return today.After(a.StartTime)
}

// IsNotEnded defines the activity is not ended
func (a *ActivityTrait) IsNotEnded() bool {
	if a.EndTime.IsZero() {
		return true
	}
	today := time.Now()
	return today.Before(a.EndTime)
}

// IsActivite defines the activity is activite
func (a *ActivityTrait) IsActivite() bool {
	return a.IsRegular() && a.IsStarted() && a.IsNotEnded()
}
