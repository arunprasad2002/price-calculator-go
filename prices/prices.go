package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxincludedPriceJob struct {
	IOManager        iomanager.IOManager `json:"-"`
	TaxRate          float64             `json:"tax_rate"`
	InputPrices      []float64           `json:"input_prices"`
	TaxIncludePrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxincludedPriceJob) Process() error {
	err := job.LoadData()

	if err != nil {
		return err
	}
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludePrices = result
	return job.IOManager.WriteResult(job)
}

func (job *TaxincludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}
	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return err
	}

	job.InputPrices = prices
	return nil
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxincludedPriceJob {
	return &TaxincludedPriceJob{
		IOManager: iom,
		TaxRate:   taxRate,
	}
}
