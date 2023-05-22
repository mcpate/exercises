package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	readInt := func() int {
		for {
			input, err := r.ReadString('\n')
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			input = strings.TrimSuffix(input, "\n")

			inputInt, err := strconv.Atoi(input)
			if err != nil {
				// If the string was not a number, try again. Otherwise, exit with failure.
				if _, ok := err.(*strconv.NumError); ok {
					fmt.Printf("%s is not an integer. Please try again: ", input)
					continue
				}
				fmt.Println(err.Error())
				os.Exit(1)
			}

			return inputInt
		}
	}

	fmt.Print("How many people? ")
	numPeople := readInt()

	fmt.Print("How many pizzas? ")
	numPizzas := readInt()

	fmt.Printf("%d people with %d pizzas.\n", numPeople, numPizzas)

	const slicesPerPizza = 8
	totalSlices := slicesPerPizza * numPizzas
	slicesPerPerson := int(math.Floor(float64(totalSlices) / float64(numPeople)))
	remainingSlices := totalSlices - (slicesPerPerson * numPeople)

	fmt.Printf("Each person gets %d slices of pizza. There are %d leftover slices.", slicesPerPerson, remainingSlices)
}
