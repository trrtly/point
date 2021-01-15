package trait

import (
	"time"
)

type (
	// Activity defines a activity is alife
	Activity struct {
		MoneyServicePoint
		StartEndTime
		NumPreDay int32  `json:"num_pre_day"`
		NumTotal  int32  `json:"num_total"`
		PointDesc string `json:"point_desc"`
	}
)

// IsStarted defines the activity is started
func (a *Activity) IsStarted() bool {
	return time.Now().After(a.StartTime)
}

// IsNotEnded defines the activity is not ended
func (a *Activity) IsNotEnded() bool {
	if a.EndTime.IsZero() {
		return true
	}
	return time.Now().Before(a.EndTime)
}

// IsActivite defines the activity is activite
func (a *Activity) IsActivite() bool {
	return a.IsRegular() && a.IsStarted() && a.IsNotEnded()
}
