package trait

import (
	"point/internal/core/status"
	"time"
)

type (
	// Activity defines a activity is alife
	Activity struct {
		Point
		StartEndTime
		NumPreDay int32  `json:"num_pre_day"`
		NumTotal  int32  `json:"num_total"`
		PointDesc string `json:"point_desc"`
	}
)

// IsRegular defines the activity record is regular
func (a *Activity) IsRegular() bool {
	return a.Status == status.Regular
}

// IsStarted defines the activity is started
func (a *Activity) IsStarted() bool {
	today := time.Now()
	return today.After(a.StartTime)
}

// IsNotEnded defines the activity is not ended
func (a *Activity) IsNotEnded() bool {
	if a.EndTime.IsZero() {
		return true
	}
	today := time.Now()
	return today.Before(a.EndTime)
}

// IsActivite defines the activity is activite
func (a *Activity) IsActivite() bool {
	return a.IsRegular() && a.IsStarted() && a.IsNotEnded()
}
