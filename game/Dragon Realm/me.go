package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	playAgain := "yes"
	for playAgain == "yes" || playAgain == "y" {
		displayIntro()
		checkCave(chooseCave())
		fmt.Println("Do you want to play again? (yes or no)")
		fmt.Scan(&playAgain)
	}
}
func displayIntro() {
	fmt.Println("You are in a land full of dragons. In front of you,")
	fmt.Println("You are in a land full of dragons. In front of you,")
	fmt.Println("and will share his treasure with you. The")
	fmt.Println("is greedy and hungry, and will eat you on sight.")
}
func chooseCave() string {
	cave := ""
	for cave != "1" && cave != "2" {
		fmt.Println("While cave will you go into? (1 or 2)")
		fmt.Scan(&cave)
	}
	return cave
}
func checkCave(chosenCave string) {
	fmt.Println("You approach the cave...")
	time.Sleep(2 * time.Second)
	fmt.Println("It is dark and spooky...")
	time.Sleep(2 * time.Second)
	fmt.Println("A large dragon jumps out in front of you! He opens his jaws and ...")
	time.Sleep(2 * time.Second)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	friendlyCave := r.Intn(2) + 1
	if chosenCave == strconv.Itoa(friendlyCave) {
		fmt.Println("Gives you his treasure")
	} else {
		fmt.Println("Gobbles you down in one bite!")
	}
}
