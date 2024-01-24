package prices

import (
	"fmt"
	"github.com/Dethblo/godethblo/internal/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) Process() error {
	// load price data from file
	err := job.LoadData()
	if err != nil {
		return err
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	err = job.IOManager.WriteResult(job)
	if err != nil {
		return err
	}

	return nil
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		TaxRate:   taxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadData() error {
	// read the data file
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	// convert lines to price values
	job.InputPrices, err = StringsToFloat(lines)
	if err != nil {
		return err
	}

	return nil
}
