package converters

type EUR Currency

var eurRates = map[string]Rate{
	"CAD": {0.68033},
	"USD": {0.82104},
}

func (e *EUR) SetAmount(amount float64) {
	e.Amount = amount
}

func (e EUR) Convert(currency string) float64 {
	rate := eurRates[currency]
	return e.Amount * rate.rate
}

func (e EUR) GetLabel() string {
	return e.Label
}

func (e EUR) GetSymbol() string {
	return e.Symbol
}
