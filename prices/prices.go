package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxincludedPriceJob struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludePrices map[string]float64
}

func (job *TaxincludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func (job *TaxincludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Could not open file")
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading the file content failed")
		fmt.Println(err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))

	for lineIndex, line := range lines {
		floatPrices, err := strconv.ParseFloat(line, 64)

		if err != nil {
			file.Close()
			fmt.Println("Converting price to float failed")
			fmt.Println(err)
		}
		prices[lineIndex] = floatPrices
	}

	job.InputPrices = prices
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxincludedPriceJob {
	return &TaxincludedPriceJob{
		TaxRate: taxRate,
	}
}
