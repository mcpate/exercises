package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
When working in a global environment, you’ll have to present information in both metric and Imperial units.
And you’ll need to know when to do the conversion to ensure the most accurate results.
Create a program that calculates the area of a room. Prompt the user for the length and width of the room in feet.
Then display the area in both square feet and square meters.

Example Output:
What is the length of the room in feet? 15
What is the width of the room in feet? 20
You entered dimensions of 15 feet by 20 feet. The area is 300 square feet, or 27.871 square meters

The formula for this conversion is: m^2 = f^2 x 0.09290304

Constraints:
- Keep the calculations separate from the output.
- Use a constant to hold the conversion factor.

*/

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("What is the length of the room in feet? ")
	length, _ := r.ReadString('\n')
	length = strings.TrimSuffix(length, "\n")

	fmt.Print("What is the width of the room in feet? ")
	width, _ := r.ReadString('\n')
	width = strings.TrimSuffix(width, "\n")

	lengthInt, _ := strconv.Atoi(length)
	widthInt, _ := strconv.Atoi(width)

	sqFeet := lengthInt * widthInt

	const sqFeetToSqMetersConversion float64 = 0.09290304
	sqMeters := float64(sqFeet) * sqFeetToSqMetersConversion

	fmt.Printf("You entered dimensions of %d feet by %d feet.\nThe area is %d square feet, or %f square meters.", lengthInt, widthInt, sqFeet, sqMeters)

}
