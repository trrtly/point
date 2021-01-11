package core

type (
	// UserAssets defines user_assets table
	UserAssets struct {
		ID           int64   `json:"id"`
		UID          int64   `json:"uid"`
		MoneyPoint   float64 `json:"money_point"`
		ServicePoint float64 `json:"service_point"`
		CreatedAt    string  `json:"-"`
		UpdatedAt    string  `json:"-"`
	}

	// UserAssetsStore defines operations for working with user_assets.
	UserAssetsStore interface {
		// Find returns a user from the datastore.
		Find(int64) (*UserAssets, error)
		// IncrPoint increment a user's money or service point or both
		IncrPoint(int64, float64, float64) error
	}
)

// TableName defines the user assets table name in db
func (UserAssets) TableName() string {
	return "t_user_assets"
}
