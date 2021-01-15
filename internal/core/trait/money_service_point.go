package trait

type (
	// MoneyServicePoint defines a money_point and service_point trait
	MoneyServicePoint struct {
		MoneyPoint   float64 `json:"money_point"`
		ServicePoint float64 `json:"service_point"`
	}
)

// IsPointGtZero defines the poin grant then zero
func (a *MoneyServicePoint) IsPointGtZero() bool {
	return a.MoneyPoint > 0 || a.ServicePoint > 0
}
