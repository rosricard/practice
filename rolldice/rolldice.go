package rolldice

import (
	"fmt"
	"math/rand"
)

// Rolldice takes an input from the user for the number of dice
// it returns the sum of the number of dice rolled

func Rolldice() (*int, error) {
	var dice int
	println("How many dice would you like to roll?")
	_, err := fmt.Scan(&dice)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return nil, err
	}

	var sum int

	// generate random number between 1 and 6 for each die and add to sum
	for i := 0; i < dice; i++ {
		randNum := rand.Intn(6)
		fmt.Println("Die", i+1, "rolled a", randNum)
		sum += randNum
		fmt.Println("Sum is now", sum)
	}

	return &sum, nil
}
