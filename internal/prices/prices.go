package prices

import (
	"fmt"
	"github.com/Dethblo/godethblo/internal/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager `json:"-"`
	TaxRate           float64                 `json:"tax_rate"`
	InputPrices       []float64               `json:"input_prices"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) Process() {
	// load price data from file
	job.LoadData()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	err := job.IOManager.WriteResult(job)
	if err != nil {
		fmt.Println(err)
	}
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: fm,
		TaxRate:   taxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadData() {
	// read the data file
	lines, err := job.IOManager.ReadLines()

	// convert lines to price values
	job.InputPrices, err = StringsToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

}
