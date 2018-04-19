package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	guessesTaken := 0
	fmt.Println("Hello!What's your name?")
	var myName string
	var guess int
	fmt.Scan(&myName)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	number := r.Intn(20)
	fmt.Println("Well," + myName + ",I am thinking of a number between 1 and 20.")
	for guessesTaken < 6 {
		fmt.Println("Take a guess.")
		fmt.Scan(&guess)
		guessesTaken = guessesTaken + 1
		if guess < number {
			fmt.Println("Your guess is to low.")
		}
		if guess > number {
			fmt.Println("Your guess is to high.")
		}
		if guess == number {
			break
		}
	}
	if guess == number {
		fmt.Printf("Good job,%v! You guessed my number in %v guesses!\n", myName, guessesTaken)
	}
	if guess != number {
		fmt.Printf("Nope. The number I was thinking of was %v", number)
	}
}
