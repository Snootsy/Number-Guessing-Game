# Number Guessing Game

**Project URL:** [(https://github.com/yourusername/number-guessing-game)](https://roadmap.sh/projects/number-guessing-game)

## Description
A simple command-line number guessing game written in Go. The computer randomly selects a number between 1 and 100, and the player tries to guess it within a limited number of attempts. The number of attempts depends on the selected difficulty level.

## Assignment Requirements
- CLI-based game.
- Display a welcome message and rules when the game starts.
- Randomly select a number between 1 and 100.
- Let the user choose a difficulty level (easy, medium, hard) that determines the number of guesses.
- Accept user guesses.
- If the guess is correct, show a congratulatory message and the number of attempts.
- If the guess is wrong, indicate whether the number is greater or less.
- End the game when the user guesses correctly or runs out of chances.

## Features
- Three difficulty levels: Easy (10 chances), Medium (5 chances), Hard (3 chances).
- Input validation for difficulty choice and guesses.
- Option to play again after finishing a round.
- Clear feedback on each guess.

## Requirements
- Go 1.20 or higher (any recent version should work).

## Installation & Run
```bash
git clone https://github.com/yourusername/number-guessing-game.git
cd number-guessing-game
go run main.go

How to Play
Run the program.

Choose difficulty: 1 (Easy), 2 (Medium), or 3 (Hard).

Enter your guess when prompted.

Receive feedback: whether the number is higher or lower.

Continue guessing until you win or run out of attempts.

After the game, choose to play again or exit.
