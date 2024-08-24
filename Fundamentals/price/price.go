package price

import (
	iomanager "example/my-app/Fundamentals/IOManager"
	"example/my-app/Fundamentals/conversion"
	"fmt"
	"time"
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

func (price *Price) Process(doneChan chan bool, errorChan chan error) {
	err := price.LoadInputPrice()

	// errorChan <- errors.New("An Error!")

	if err != nil {
		fmt.Println(err)
		errorChan <- err
	}

	result := make(map[string]float64)
	for _, priceRate := range price.InputPrice {
		taxIncludedPrice := priceRate * (1 + price.TaxRate/100)
		result[fmt.Sprintf("%0.2f", priceRate)] = taxIncludedPrice
	}
	price.TaxIncludedPriceMap = result
	fmt.Println(*price)

	time.Sleep(3 * time.Second)
	err = price.IOManager.WriteJSON(price)
	if err != nil {
		fmt.Println(err)
		errorChan <- err
	}
	doneChan <- true
}

func NewPrice(taxRate float64, iOManager iomanager.IOManager) *Price {
	return &Price{
		IOManager:  iOManager,
		TaxRate:    taxRate,
		InputPrice: []float64{8.99, 11.2, 4.5},
	}
}
