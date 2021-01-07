package core

import (
	"time"
)

type (
	// ActivitySpecial defines user_assets table
	ActivitySpecial struct {
		ID            int32     `json:"id"`
		ActivityID    int32     `json:"activity_id"`
		Name          string    `json:"name"`
		SType         string    `json:"s_type"`
		SValue        string    `json:"s_value"`
		STableName    string    `json:"s_table_name"`
		StartTime     time.Time `json:"start_time"`
		EndTime       time.Time `json:"end_time"`
		MoneyPoint    float64   `json:"money_point"`
		ServicePoint  float64   `json:"service_point"`
		NumPreDay     int32     `json:"num_pre_day"`
		NumTotal      int32     `json:"num_total"`
		CreatedUserID string    `json:"created_user_id"`
		Status        int8      `json:"status"`
		CreatedTime   string    `json:"-"`
		ModifyTime    string    `json:"-"`
	}

	// ActivitySpecialStore defines operations for working with user_assets.
	ActivitySpecialStore interface {
		// FindSVal returns a special val activity from the db.
		FindSVal(int32, string) (*ActivitySpecial, error)
	}
)

// TableName defines the activity table name in db
func (ActivitySpecial) TableName() string {
	return "t_activity_special"
}
