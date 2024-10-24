package main

import (
	"fmt"
	"go-challange/blackjack/cards"
)

func main() {
	var card cards.Card
	var sum int
	deck := cards.CreateDeck()
	var input string
	fmt.Println("BLACK JACK")
	for sum < 21 {

		fmt.Println("DRAW A CARD! - (y) to drawn"
		fmt.Scan(&input)
		card = *cards.DrawCard(deck)
		sum += int(card.Points)
		fmt.Println(sum)
	}

}
