package special

import (
	"point/internal/core"
	"point/internal/store/shared/db"
)

// New returns a new ActivitySpecialStore.
func New(db *db.DB) core.ActivitySpecialStore {
	return &specialStore{db}
}

type specialStore struct {
	db *db.DB
}

// FindSVal returns a activity special from the datastore.
func (s *specialStore) FindSVal(
	activityID int32, sval string,
) (*core.ActivitySpecial, error) {
	out := &core.ActivitySpecial{}
	err := s.db.Select("*").
		Where("activity_id = ?", activityID).
		Where("s_value = ?", sval).
		First(out).Error
	return out, err
}
