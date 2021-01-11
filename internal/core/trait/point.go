package trait

type (
	// Point defines a activity is alife
	Point struct {
		MoneyPoint   float64 `json:"money_point"`
		ServicePoint float64 `json:"service_point"`
	}
)

// IsPointGtZero defines the poin grant then zero
func (a *Point) IsPointGtZero() bool {
	return a.MoneyPoint > 0 || a.ServicePoint > 0
}
