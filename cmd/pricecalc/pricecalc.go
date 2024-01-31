package main

import (
	"fmt"
	"github.com/Dethblo/godethblo/internal/filemanager"
	"github.com/Dethblo/godethblo/internal/prices"
	"sync"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	var wg sync.WaitGroup

	for _, taxRate := range taxRates {
		mgr := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))

		// the following mgr line can be uncommented and the above line commented to switch between
		// using the Console vs Files (illustrating the usage of interfaces in design)
		//mgr := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(mgr, taxRate)
		wg.Add(1)
		go priceJob.Process(&wg)

		//if err != nil {
		//	fmt.Println("Could not process Job")
		//	fmt.Println(err)
		//}

	}

	// wait for all goroutines to complete
	wg.Wait()
}
