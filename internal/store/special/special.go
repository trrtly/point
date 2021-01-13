package special

import (
	"point/internal/core"
	"point/internal/store/shared/db"
)

// New returns a new ActivitySpecialStore.
func New(d *db.DB) core.ActivitySpecialStore {
	m := d.Model(&core.ActivitySpecial{})
	return &specialStore{&db.DB{DB: m}}
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
