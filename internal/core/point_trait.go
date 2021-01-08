package core

type (
	// PointTrait defines a activity is alife
	PointTrait struct {
		MoneyPoint   float64 `json:"money_point"`
		ServicePoint float64 `json:"service_point"`
	}
)

// IsPointGtZero defines the poin grant then zero
func (a *PointTrait) IsPointGtZero() bool {
	return a.MoneyPoint > 0 || a.ServicePoint > 0
}
