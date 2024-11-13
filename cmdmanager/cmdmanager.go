package cmdmanager

import "fmt"

type CMDmanager struct {
}

func (cmd CMDmanager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER")
	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		prices = append(prices, price)

		if price == "0" {
			break
		}
	}

	return prices, nil
}

func (cmd CMDmanager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDmanager {
	return CMDmanager{}
}
