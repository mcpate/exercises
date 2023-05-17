package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Challenge:
Quotation marks are often used to denote the start and end of a string. But sometimes
we need to print out the quotation marks themselves by using escape characters.
Create a program that prompts for a quote and an author. Display the quotation and author
as shown in the example output.

Example Output: What is the quote? These aren't the droids you're looking for.
Who said it? Obi-Wan Kenobi
Obi-Wan Kenobi says, "These aren't the droids you're looking for."

Constraints:
- Use a single output statement to produce this output, using appropriate string-escaping techniques for quotes.
- If your language supports string interpolation or string substitution, don’t use it for this exercise. Use string concatenation instead.

My notes:
- No requirements for blank entries OR quoted strings, so assuming fine.
*/

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("What is the quote? ")
	quote, _ := readString(r)

	fmt.Print("Who said it? ")
	author, _ := readString(r)

	fmt.Println(author + " says, \"" + quote + "\"")
}

func readString(r *bufio.Reader) (string, error) {
	str, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(str, "\n"), nil
}
