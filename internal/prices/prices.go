package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return
	}

	// make sure file is closed when done
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Could not close file!")
			fmt.Println(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	// for each line in the file (until EOF returns false)...
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// check for any error emitted by the scanner
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading the file!")
		fmt.Println(err)
		return
	}

	// convert lines to price values
	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, conErr := strconv.ParseFloat(line, 64)
		if conErr != nil {
			fmt.Printf("Converting string to float failed for %v", line)
			fmt.Println(err)
			return
		}

		prices[lineIndex] = floatPrice
	}

	job.InputPrices = prices
}
