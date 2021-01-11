package core

import (
	"point/internal/core/status"
	"point/internal/core/trait"
)

type (
	// Activity defines user_assets table
	Activity struct {
		ID            int32  `json:"id"`
		Yyid          string `json:"yyid"`
		Name          string `json:"name"`
		EventKey      string `json:"event_key"`
		ExistSpecial  int8   `json:"exist_special"`
		CreatedUserID string `json:"created_user_id"`
		CreatedTime   string `json:"-"`
		ModifyTime    string `json:"-"`
		trait.Activity
	}

	// ActivityStore defines operations for working with user_assets.
	ActivityStore interface {
		// FindEventKey returns a special event_key activity from the datastore.
		FindEventKey(string) (*Activity, error)
	}
)

const (
	// ActivityTypeGain 发放
	ActivityTypeGain = 1
	// ActivityTypeUse 使用
	ActivityTypeUse = 2
)

// TableName defines the activity table name in db
func (Activity) TableName() string {
	return "t_activity"
}

// HasSpecial defines the activity has special
func (a *Activity) HasSpecial() bool {
	return a.ExistSpecial == status.Regular
}
