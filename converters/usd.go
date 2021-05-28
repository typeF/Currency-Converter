package converters

type USD Currency

var usdRates = map[string]Rate{
	"CAD": {0.82861},
	"EUR": {1.2180},
}

func (u *USD) SetAmount(amount float64) {
	u.Amount = amount
}

func (u USD) Convert(currency string) float64 {
	rate := usdRates[currency]
	return u.Amount * rate.rate
}

func (u USD) GetLabel() string {
	return u.Label
}

func (u USD) GetSymbol() string {
	return u.Symbol
}




