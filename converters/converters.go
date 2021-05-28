package converters

type Converter interface {
	SetAmount(amount float64)
	Convert(currency string) float64
	GetLabel() string
	GetSymbol() string
}

type Currency struct {
	Label string
	Symbol string
	Amount float64
}

type Rate struct {
	rate float64
}