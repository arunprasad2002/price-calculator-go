package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxincludedPriceJob struct {
	IOManager        filemanager.FileManager `json:"-"`
	TaxRate          float64                 `json:"tax_rate"`
	InputPrices      []float64               `json:"input_prices"`
	TaxIncludePrices map[string]string       `json:"tax_included_prices"`
}

func (job *TaxincludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludePrices = result
	job.IOManager.WriteResult(job)
}

func (job *TaxincludedPriceJob) LoadData() {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices

}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxincludedPriceJob {
	return &TaxincludedPriceJob{
		IOManager: fm,
		TaxRate:   taxRate,
	}
}
