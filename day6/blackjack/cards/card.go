package cards

import "math/rand"

type Card struct {
	Suit   string
	Points uint8
	Face   string
}

func DrawCard(cards []Card) *Card {
	idx := rand.Intn(len(cards) - 1)
	card := cards[idx]
	return &card
}

func CreateDeck() []Card {

	cardAmount := 52
	suitsAmount := cardAmount / 13
	var result []Card

	for i := 0; i < suitsAmount; i++ {
		suit := CreateFullSuit(uint8(i))
		result = append(result, suit...)
	}
	return result

}

func CreateFullSuit(key uint8) []Card {
	amount := 13
	var result []Card
	for i := 1; i <= amount; i++ {
		card := CreateCard(key, uint8(i))
		result = append(result, *card)
	}

	return result

}
func CreateCard(suit, points uint8) *Card {
	card := Card{
		Suit:   getSuitsMap(suit),
		Points: uint8(points),
		Face:   "",
	}

	if points == 1 {
		card.Face = "AS"
	}
	if points > 10 {
		card.Points = 10
	}
	if points == 11 {
		card.Face = "Jack"
	} else if points == 12 {
		card.Face = "Queen"
	} else if points >= 13 {
		card.Face = "King"
	}

	return &card
}

func getSuitsMap(key uint8) string {

	suitsMap := make(map[uint8]string)
	suitsMap[0] = "Hearts"
	suitsMap[1] = "Clubs"
	suitsMap[2] = "Diamonds"
	suitsMap[3] = "Spades"
	return suitsMap[key]
}
