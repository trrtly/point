package activity

import (
	"point/internal/core"
	"point/internal/store/shared/db"
)

// New returns a new UserStore.
func New(d *db.DB) core.ActivityStore {
	m := d.Model(&core.Activity{})
	return &activityStore{&db.DB{DB: m}}
}

type activityStore struct {
	db *db.DB
}

// FindEventKey returns a activity from the datastore.
func (s *activityStore) FindEventKey(eventKey string) (*core.Activity, error) {
	out := &core.Activity{}
	err := s.db.Select("*").
		Where("event_key = ?", eventKey).
		First(out).Error
	return out, err
}
