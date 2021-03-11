package engine

import (
	"fmt"
)

type Card struct {
	Type  CardType `json:"type"`
	Color Color    `json:"color"`
	Num   uint8    `json:"numberValue"`
}

var allCards = map[Card]uint8{}

func GenerateAllCards() {
	for color := Red; color <= Yellow; color++ {
		for i := 0; i < 10; i++ {
			if i != 0 {
				allCards[Card{Color: color, Num: uint8(i), Type: Numeric}] = 2
				continue
			}
			allCards[Card{Color: color, Num: uint8(i), Type: Numeric}] = 1
		}
		allCards[Card{Color: color, Type: Reverse}] = 2
		allCards[Card{Color: color, Type: Skip}] = 2
		allCards[Card{Color: color, Type: TakeTwo}] = 2
	}
	allCards[Card{Type: ChooseColor}] = 4
	allCards[Card{Type: TakeFourChooseColor}] = 4
	fmt.Println("All card added to map")
}

func GetOpponentHand(hand []Card, discard []Card) map[Card]uint8 {
	r := make(map[Card]uint8, len(allCards))
	for k := range allCards {
		r[k] = allCards[k]
	}

	for i := 0; i < len(hand); i++ {
		if r[hand[i]] > 0 {
			r[hand[i]]--
		}
	}

	for i := 0; i < len(discard); i++ {
		if r[discard[i]] > 0 {
			r[discard[i]]--
		}
	}
	return r
}
