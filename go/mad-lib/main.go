package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Mad libs are a simple game where you create a story template with blanks for words. You, or another player,
then construct a list of words and place them into the story, creating an often silly or funny story as a result.
Create a simple mad-lib program that prompts for a noun, a verb, an adverb, and an adjective and injects those into a story that you create.

Example Output:
Enter a noun: dog
Enter a verb: walk
Enter an adjective: blue
Enter an adverb: quickly
Do you walk your blue dog quickly? That's hilarious!

Constraints:
- Use a single output statement for this program.
- If your language supports string interpolation or string substitution, use it to build up the output.
*/

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a noun: ")
	noun1, _ := read(r)

	fmt.Print("Enter a noun: ")
	noun2, _ := read(r)

	fmt.Print("Enter a noun representing a sound: ")
	noun3, _ := read(r)

	fmt.Print("Enter an adverb: ")
	adverb1, _ := read(r)

	story := "The %s sat on the %s and made a loud %s. He %s got to his feet and ran away."

	fmt.Println(fmt.Sprintf(story, noun1, noun2, noun3, adverb1))

}

func read(r *bufio.Reader) (string, error) {
	str, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(str, "\n"), nil
}
