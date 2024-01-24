package prices

import (
	"fmt"
	"github.com/Dethblo/godethblo/internal/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job TaxIncludedPriceJob) Process() {
	// load price data from file
	job.LoadData()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadData() {
	// read the data file
	lines, err := filemanager.ReadLines("prices.txt")

	// convert lines to price values
	job.InputPrices, err = StringsToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

}
