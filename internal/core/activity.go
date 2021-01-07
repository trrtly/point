package core

import "time"

type (
	// Activity defines user_assets table
	Activity struct {
		ID            int32   `json:"id"`
		Yyid          string  `json:"yyid"`
		Name          string  `json:"name"`
		EventKey      string  `json:"event_key"`
		StartTime     string  `json:"start_time"`
		EndTime       string  `json:"end_time"`
		MoneyPoint    float64 `json:"money_point"`
		ServicePoint  float64 `json:"service_point"`
		ExistSpecial  int8    `json:"exist_special"`
		NumPreDay     int32   `json:"num_pre_day"`
		NumTotal      int32   `json:"num_total"`
		CreatedUserID string  `json:"created_user_id"`
		Status        int8    `json:"status"`
		CreatedTime   string  `json:"-"`
		ModifyTime    string  `json:"-"`
	}

	// ActivityStore defines operations for working with user_assets.
	ActivityStore interface {
		// FindEventKey returns a special event_key activity from the datastore.
		FindEventKey(string) (*Activity, error)
	}
)

// TableName defines the activity table name in db
func (Activity) TableName() string {
	return "t_activity"
}

// IsRegular defines the activity record is regular
func (a *Activity) IsRegular() bool {
	return a.Status == StatusRegular
}

// IsStarted defines the activity is started
func (a *Activity) IsStarted() bool {
	today := time.Now()
	startTime, err := time.Parse("2006-01-02 03:04:05", a.StartTime)
	if err != nil {
		return false
	}
	return today.After(startTime)
}

// IsNotEnded defines the activity is not ended
func (a *Activity) IsNotEnded() bool {
	today := time.Now()
	endTime, err := time.Parse("2006-01-02 03:04:05", a.EndTime)
	if err != nil {
		return false
	}
	return today.Before(endTime)
}

// IsActivite defines the activity is activite
func (a *Activity) IsActivite() bool {
	return a.IsRegular() && a.IsStarted() && a.IsNotEnded()
}
