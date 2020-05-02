package main

import "fmt"

func main() {
	//
	// The variable declaration
	//
	// var card string = "Ace of Spades"
	//
	// is equivalent to
	//
	// card := "Ace of Spades"
	//
	card := newCard()
	fmt.Println(card)
}

/*
func newCard() string {
	return "Ace of Spades"
}
*/
