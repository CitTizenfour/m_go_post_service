package code_generator

import (
	"math/rand"
	"time"
)

func GetCode() string {
	// Set the seed value for the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a string of uppercase letters and numbers
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Initialize counters for letters and numbers
	numCount := 0
	letterCount := 0

	// Initialize an empty string to hold the result
	result := ""

	// Generate a random string with a length of 6
	for i := 0; i < 6; i++ {
		// Generate a random index in the range of the available characters
		index := rand.Intn(len(chars))

		// Get the character at the randomly generated index
		char := chars[index]

		// Check if the character is a number or a letter
		if char >= '0' && char <= '9' {
			// If the character is a number, check if the maximum number count has been reached
			if numCount >= 3 {
				// If the maximum number count has been reached, generate a letter instead
				index = rand.Intn(26)
				char = chars[index]
			} else {
				// If the maximum number count has not been reached, increment the number count
				numCount++
			}
		} else {
			// If the character is a letter, check if the maximum letter count has been reached
			if letterCount >= 3 {
				// If the maximum letter count has been reached, generate a number instead
				index = rand.Intn(10) + 26
				char = chars[index]
			} else {
				// If the maximum letter count has not been reached, increment the letter count
				letterCount++
			}
		}

		// Add the generated character to the result string
		result += string(char)
	}

	return result
}
