package core

import (
	"time"
)

type (
	// Activity defines user_assets table
	Activity struct {
		ID            int32     `json:"id"`
		Yyid          string    `json:"yyid"`
		Name          string    `json:"name"`
		EventKey      string    `json:"event_key"`
		StartTime     time.Time `json:"start_time"`
		EndTime       time.Time `json:"end_time"`
		MoneyPoint    float64   `json:"money_point"`
		ServicePoint  float64   `json:"service_point"`
		ExistSpecial  int8      `json:"exist_special"`
		NumPreDay     int32     `json:"num_pre_day"`
		NumTotal      int32     `json:"num_total"`
		CreatedUserID string    `json:"created_user_id"`
		Status        int8      `json:"status"`
		CreatedTime   string    `json:"-"`
		ModifyTime    string    `json:"-"`
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

// HasSpecial defines the activity has special
func (a *Activity) HasSpecial() bool {
	return a.ExistSpecial == StatusRegular
}
