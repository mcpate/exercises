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

	readFloat := func() float64 {
		for {
			input, err := r.ReadString('\n')
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			input = strings.TrimSuffix(input, "\n")

			inputF, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			return inputF
		}
	}

	fmt.Print("What is the length of the room in feet? ")
	length := readFloat()

	fmt.Print("What is the width of the room in feet? ")
	width := readFloat()

	const sqFtPerGallon = 350.0

	totalSqFt := length * width

	numGallons := int(math.Ceil(totalSqFt / sqFtPerGallon))

	fmt.Printf("You will need to purchase %d gallons of paint to cover %.3f square feet", numGallons, totalSqFt)
}
