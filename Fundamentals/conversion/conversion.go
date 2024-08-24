package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloats(strs []string) ([]float64, error) {
	convertedFloats := make([]float64, len(strs))

	for index, str := range strs {
		floatVal, err := strconv.ParseFloat(str, 64)

		if err != nil {
			fmt.Println("Error converting string to float")
			return nil, err
		}

		convertedFloats[index] = floatVal
	}

	return convertedFloats, nil
}
