package prices

import (
	"errors"
	"fmt"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	var prices []float64
	for _, stringVal := range strings {
		floatPrice, conErr := strconv.ParseFloat(stringVal, 64)
		if conErr != nil {
			return nil, errors.New(fmt.Sprintf("Failed to convert string: [%v]", stringVal))
		}

		prices = append(prices, floatPrice)
	}

	return prices, nil
}
