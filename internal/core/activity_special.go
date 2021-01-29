package core

import "point/internal/core/trait"

type (
	// ActivitySpecial defines user_assets table
	ActivitySpecial struct {
		ID            int32  `json:"id"`
		ActivityID    int32  `json:"activity_id"`
		Name          string `json:"name"`
		SType         string `json:"s_type"`
		SValue        string `json:"s_value"`
		STableName    string `json:"s_table_name"`
		CreatedUserID string `json:"created_user_id"`
		CreatedTime   string `json:"-"`
		ModifyTime    string `json:"-"`
		trait.Activity
	}

	// ActivitySpecialStore defines operations for working with user_assets.
	ActivitySpecialStore interface {
		// FindSVal returns a special val activity from the db.
		FindSVal(int32, string) (*ActivitySpecial, error)
	}
)

// TableName defines the table name
func (ActivitySpecial) TableName() string {
	return "t_activity_special"
}
