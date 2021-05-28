package main

import (
	"CurrencyConverter/converters"
	"CurrencyConverter/validate"
	"fmt"
	"strconv"
)

func main() {
	c := &converters.CAD{Label: "CAD", Symbol: "$", Amount: 0}
	u := &converters.USD{Label: "USD", Symbol: "$", Amount: 0}
	e := &converters.EUR{Label: "EUR", Symbol: "€", Amount: 0}
	converters := []converters.Converter{c, u, e}

	fmt.Println("Select your base currency:")
	for i, v := range converters {
		fmt.Printf("[%v] %v\n", i+1, v.GetLabel())
	}

	var baseCurrency int
	fmt.Scanln(&baseCurrency)

	for validate.OutsideSelectionRange(baseCurrency, len(converters)) {
		fmt.Println("Invalid selection, try again:")
		fmt.Scanln(&baseCurrency)
	}

	selectedCurrency := converters[baseCurrency-1]
	currencyLabel := selectedCurrency.GetLabel()
	currencySymbol := selectedCurrency.GetSymbol()

	fmt.Printf("Selected base currency: %v\n\n", currencyLabel)
	fmt.Println("Select your currency to convert to:")
	for i := 0; i < len(converters); i++ {
		if i != (baseCurrency - 1) {
			fmt.Printf("[%v] %v\n", i+1, converters[i].GetLabel())
		}
	}

	var selectedConverter int
	fmt.Scanln(&selectedConverter)

	for selectedConverter == baseCurrency || selectedConverter < 1 || selectedConverter > len(converters) {
		fmt.Println("Invalid selection, try again:")
		fmt.Scanln(&selectedConverter)
	}

	converter := converters[selectedConverter-1]
	converterLabel := converter.GetLabel()
	converterSymbol := converter.GetSymbol()

	fmt.Printf("Selected conversion currency: %v\n\n", converterLabel)

	fmt.Printf("Enter %v amount to convert:\n", currencyLabel)
	fmt.Printf("%v", currencySymbol)
	var amount float64
	fmt.Scanln(&amount)

	converter.SetAmount(amount)
	convertedAmount := converter.Convert(currencyLabel)
	fmt.Printf("%v%v %v -> %v%v %v\n", currencySymbol, amount, currencyLabel, converterSymbol, strconv.FormatFloat(convertedAmount, 'f', 2, 64), converterLabel)
}
