package main

import (
	"bufio"
	"fmt"
	"github.com/shopspring/decimal"
	"os"
	"strings"
)

/*
You’ll often write programs that deal with numbers. And depending
on the programming language you use, you’ll have to convert the
inputs you get to numerical data types.
Write a program that prompts for two numbers. Print the sum, difference,
product, and quotient of those numbers as shown in the example output.

Example Output:
What is the first number? 10
What is the second number? 5
10 + 5 = 15
10 - 5 = 5
10 * 5 = 50
10 / 5 = 2

Constraints:
- Values coming from users will be strings. Ensure that you convert these values to numbers before
doing the math.
- Keep the inputs and outputs separate from the numerical conversions and other processing.
- Generate a single output statement with line breaks in the appropriate spots.
*/
func main() {
	fmt.Print("What is the first number? ")
	r := bufio.NewReader(os.Stdin)
	left, _ := r.ReadString('\n')
	left = strings.TrimSuffix(left, "\n")
	leftD, _ := decimal.NewFromString(left)

	fmt.Print("What is the second number? ")
	right, _ := r.ReadString('\n')
	right = strings.TrimSuffix(right, "\n")
	rightD, _ := decimal.NewFromString(right)

	add := leftD.Add(rightD).String()
	sub := leftD.Sub(rightD).String()
	mul := leftD.Mul(rightD).String()
	div := leftD.Div(rightD).String()

	fmt.Println(
		left + " + " + right + " = " + add + "\n" +
			left + " - " + right + " = " + sub + "\n" +
			left + " * " + right + " = " + mul + "\n" +
			left + " / " + right + " = " + div,
	)
}
