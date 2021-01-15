package trait

import (
	"point/internal/core/status"
)

type (
	// Status defines a Status is alife
	Status struct {
		Status int8 `json:"-"`
	}
)

// IsRegular defines the Status record is regular
func (a *Status) IsRegular() bool {
	return a.Status == status.Regular
}
