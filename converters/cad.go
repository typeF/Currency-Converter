package converters

type CAD Currency

var cadRates = map[string]Rate{
	"USD": {1.2068},
	"EUR": {1.4699},
}

func (c *CAD) SetAmount(amount float64) {
	c.Amount = amount
}

func (c CAD) Convert(currency string) float64 {
	rate := cadRates[currency]
	return c.Amount * rate.rate
}

func (c CAD) GetLabel() string {
	return c.Label
}

func (c CAD) GetSymbol() string {
	return c.Symbol
}
