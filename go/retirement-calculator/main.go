package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
Your computer knows what the current year is, which means you can incorporate
that into your programs. You just have to figure out how your
programming language can provide you with that information.
Create a program that determines how many years you have left until retirement
and the year you can retire. It should prompt for your current age and the
age you want to retire and display the output as shown in the example that follows.

Example Output:
What is your current age? 25
At what age would you like to retire? 65
You have 40 years left until you can retire. It's 2015, so you can retire in 2055.

Constraints:
- Again, be sure to convert the input to numerical data before doing any math.
- Don’t hard-code the current year into your program. Get it from the system time via your programming language.
*/
func main() {
	fmt.Print("What is your current currentAge? ")
	r := bufio.NewReader(os.Stdin)

	currentAge, _ := r.ReadString('\n')
	currentAge = strings.TrimSuffix(currentAge, "\n")
	currentAgeI, err := strconv.Atoi(currentAge)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print("At what age would you like to retire? ")
	retirementAge, _ := r.ReadString('\n')
	retirementAge = strings.TrimSuffix(retirementAge, "\n")
	retirementAgeI, _ := strconv.Atoi(retirementAge)

	yearsUntilRetirement := retirementAgeI - currentAgeI
	currentYear := time.Now().Year()
	retirementYear := currentYear + yearsUntilRetirement

	fmt.Printf("You have %d years left until you can retire.\nIt is %d, so you can retire in %d.\n", yearsUntilRetirement, currentYear, retirementYear)
}
