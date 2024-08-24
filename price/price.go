package price

import (
	iomanager "example/my-app/IOManager"
	"example/my-app/conversion"
	"fmt"
)

type Price struct {
	IOManager           iomanager.IOManager `json:"-"`
	TaxRate             float64             `json:"tax_rate"`
	InputPrice          []float64           `json:"input_price"`
	TaxIncludedPriceMap map[string]float64  `json:"tax_included_price"`
}

func (price *Price) LoadInputPrice() error {
	lines, err := price.IOManager.ReadLines()
	if err != nil {
		fmt.Println("There was an error converting string array to float array")
		fmt.Println(err)
		return err
	}

	inputData, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("There was an error converting string array to float array")
		fmt.Println(err)
		return err
	}

	price.InputPrice = inputData
	return nil
}

func (price *Price) Process() {
	price.LoadInputPrice()

	result := make(map[string]float64)
	for _, priceRate := range price.InputPrice {
		taxIncludedPrice := priceRate * (1 + price.TaxRate/100)
		result[fmt.Sprintf("%0.2f", priceRate)] = taxIncludedPrice
	}
	price.TaxIncludedPriceMap = result
	fmt.Println(*price)

	price.IOManager.WriteJSON(price)
}

func NewPrice(taxRate float64, iOManager iomanager.IOManager) *Price {
	return &Price{
		IOManager:  iOManager,
		TaxRate:    taxRate,
		InputPrice: []float64{8.99, 11.2, 4.5},
	}
}
