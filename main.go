package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) 
	var playAgain string
	var highScores = map[string]int{
		"easy":   0,
		"medium": 0,
		"hard":   0,
	}

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")

	for {
		var maxAttempts int
		var difficulty string

		fmt.Println("Please select the difficulty level:")
		fmt.Println("1. Easy (10 chances)")
		fmt.Println("2. Medium (5 chances)")
		fmt.Println("3. Hard (3 chances)")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			difficulty = "easy"
			maxAttempts = 10
		case 2:
			difficulty = "medium"
			maxAttempts = 5
		case 3:
			difficulty = "hard"
			maxAttempts = 3
		default:
			fmt.Println("Invalid choice. Please select a valid difficulty level.")
			continue
		}

		fmt.Printf("Great! You have selected the %s difficulty level.\n", difficulty)
		targetNumber := rand.Intn(100) + 1 // Generate a random number between 1 and 100
		attempts := 0
		var guess int
		startTime := time.Now()

		for attempts < maxAttempts {
			fmt.Print("Enter your guess: ")
			fmt.Scan(&guess)
			attempts++

			if guess < targetNumber {
				fmt.Println("Incorrect! The number is greater than", guess)
			} else if guess > targetNumber {
				fmt.Println("Incorrect! The number is less than", guess)
			} else {
				elapsedTime := time.Since(startTime)
				fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", attempts)
				fmt.Printf("Time taken: %s\n", elapsedTime)
				if highScores[difficulty] == 0 || attempts < highScores[difficulty] {
					highScores[difficulty] = attempts
					fmt.Printf("New high score for %s difficulty: %d attempts!\n", difficulty, attempts)
				}
				break
			}
		}

		if attempts == maxAttempts {
			fmt.Printf("Sorry, you've run out of chances! The correct number was %d.\n", targetNumber)
		}

		fmt.Print("Do you want to play again? (yes/no): ")
		fmt.Scan(&playAgain)
		if playAgain != "yes" {
			break
		}
	}

	fmt.Println("Thanks for playing! Here are your high scores:")
	for level, score := range highScores {
		if score > 0 {
			fmt.Printf("%s: %d attempts\n", level, score)
		} else {
			fmt.Printf("%s: No high score yet\n", level)
		}
	}
}
