package main

import (
	"bufio"
	"fmt"
	"github.com/shopspring/decimal"
	"os"
	"strings"
)

type spec struct {
	price    decimal.Decimal
	quantity decimal.Decimal
}

func main() {
	r := bufio.NewReader(os.Stdin)

	logAndExitOnErr := func(err error) {
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	specs := make([]spec, 3)
	for i := 1; i <= 3; i++ {
		fmt.Printf("Enter the price of item %d: ", i)

		price, err := r.ReadString('\n')
		logAndExitOnErr(err)

		priceD, err := decimal.NewFromString(strings.TrimSuffix(price, "\n"))
		logAndExitOnErr(err)

		fmt.Printf("Enter the quantity of item %d: ", i)

		quantity, err := r.ReadString('\n')
		logAndExitOnErr(err)

		quantityD, err := decimal.NewFromString(strings.TrimSuffix(quantity, "\n"))
		logAndExitOnErr(err)

		specs[i-1] = spec{priceD, quantityD}
	}

	subtotal := decimal.NewFromInt(0)
	for _, spec := range specs {
		subtotal = subtotal.Add(spec.price.Mul(spec.quantity))
	}
	// 5.5% tax
	tax := decimal.NewFromFloat(.055).Mul(subtotal)
	total := subtotal.Add(tax)

	fmt.Printf("Subtotal: $%s\nTax: $%s\nTotal: $%s\n", subtotal.Truncate(2).String(), tax.Truncate(2).String(), total.Truncate(2).String())
}
