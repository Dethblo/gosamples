package main

import (
	"fmt"
	"github.com/Dethblo/godethblo/internal/filemanager"
	"github.com/Dethblo/godethblo/internal/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		mgr := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))

		// the following mgr line can be uncommented and the above line commented to switch between
		// using the Console vs Files (illustrating the usage of interfaces in design)
		//mgr := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(mgr, taxRate)
		priceJob.Process()
	}
}
