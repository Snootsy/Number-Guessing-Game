package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have several attempts to guess the correct number.")
	for {
		playAgain := playGame(reader)
		if !playAgain {
			break
		}
	}

	fmt.Println("Press Enter to exit...")
	reader.ReadString('\n')
}

func playGame(reader *bufio.Reader) bool {
	target := rand.Intn(100) + 1

	var maxGuesses int
	var difficultyName string

	for {
		fmt.Println("")
		fmt.Println("Please select the difficulty level:")
		fmt.Println("1. Easy (10 chances)")
		fmt.Println("2. Medium (5 chances)")
		fmt.Println("3. Hard (3 chances)")
		fmt.Println("")
		fmt.Print("Enter your choice:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter 1, 2, or 3.")
			continue
		}

		switch choice {
		case 1:
			maxGuesses = 10
			difficultyName = "Easy"
		case 2:
			maxGuesses = 5
			difficultyName = "Medium"
		case 3:
			maxGuesses = 3
			difficultyName = "Hard"
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
			continue
		}
		break
	}

	fmt.Printf("Great! You have selected the %s difficulty level.\n", difficultyName)
	fmt.Println("Let's start the game!")
	fmt.Println("")

	success := false
	attempts := 0

	for guesses := 0; guesses < maxGuesses; guesses++ {
		attempts = guesses + 1
		fmt.Print("Enter your guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			guesses--
			continue
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			guesses--
			continue
		}

		if guess < target {
			fmt.Printf("Incorrect! The number is greater than %d.\n", guess)
		} else if guess > target {
			fmt.Printf("Incorrect! The number is less than %d.\n", guess)
		} else {
			success = true
			fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempts)
			break
		}
	}

	if !success {
		fmt.Printf("Sorry, you've run out of attempts. The number was %d.\n", target)
	}

	for {
		fmt.Print("Play again? (y/n): ")
		againInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}
		againInput = strings.TrimSpace(strings.ToLower(againInput))

		if againInput == "y" {
			return true
		} else if againInput == "n" {
			return false
		} else {
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}
	}
}
