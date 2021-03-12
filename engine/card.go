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

func GetOpponentHand(hand []Card, discard []Card) []Card {
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

	rs := make([]Card, 0)
	for k, v := range r {
		for i := 1; i <= int(v); i++ {
			rs = append(rs, k)
		}
	}

	return rs
}

func CanNextMove(c *Card, d *Card) bool {
	if d.Type == Skip || d.Type == TakeTwo || d.Type == TakeFourChooseColor {
		return false
	}
	if c.Type == ChooseColor || c.Type == TakeFourChooseColor {
		return true
	}
	if c.Color == d.Color || c.Num == d.Num {
		return true
	}
	return false
}

func (c *Card) String() string {
	return fmt.Sprintf("Color = %s, Type = %s, Number = %d", c.color(), c.cType(), c.Num)
}

func (c *Card) color() string  {
	switch c.Color {
	case  Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	default:
		return "Black"
	}
}

func (c *Card) cType() string  {
	switch c.Type {
	case  Numeric:
		return "Numeric"
	case Reverse:
		return "Reverse"
	case Skip:
		return "Skip"
	case  TakeTwo:
		return "TakeTwo"
	case TakeFourChooseColor:
		return "TakeFourChooseColor"
	case ChooseColor:
		return "ChooseColor"
	default:
		return "????"
	}
}
