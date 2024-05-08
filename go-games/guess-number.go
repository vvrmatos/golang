package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Hardcoded secret number
	secretNumber := 5

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Guess the number I am thinking of. It's between 1 and 10.")

	for {
		fmt.Print("Enter your guess: ")
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)  // Trim any whitespace from the input
		guess, err := strconv.Atoi(input) // Convert string to integer

		if err != nil {
			fmt.Println("Please enter a valid number!")
			continue
		}

		if guess < secretNumber {
			fmt.Println("Too low!")
		} else if guess > secretNumber {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Correct! You've guessed the number!")
			break
		}
	}
}
