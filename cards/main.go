package main

import "fmt"

func main() {
	cards := []string{"Six of Diamonds", newCard()}

	cards = append(cards, "King of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
