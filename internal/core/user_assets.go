package core

import "time"

type (
	// UserAssets defines user_assets table
	UserAssets struct {
		UID          int64     `json:"uid" gorm:"primaryKey"`
		MoneyPoint   float64   `json:"money_point"`
		ServicePoint float64   `json:"service_point"`
		CreatedAt    time.Time `json:"-"`
		UpdatedAt    time.Time `json:"-"`
	}

	// UserAssetsStore defines operations for working with user_assets.
	UserAssetsStore interface {
		// Find returns a user from the datastore.
		Find(int64) (*UserAssets, error)
		// FindOrCreate create a new user assets record if not exist.
		FindOrCreate(int64) (*UserAssets, error)
		// IncrPoint increment a user's money or service point or both
		IncrPoint(int64, float64, float64) error
		// DecrPoint decrement a user's money or service point or both
		DecrPoint(int64, *UserAssets, *UserAssets) error
	}
)

// TableName defines the user assets table name in db
func (UserAssets) TableName() string {
	return "t_user_assets"
}
