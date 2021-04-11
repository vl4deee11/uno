package engine

import (
	"fmt"
)

type Card struct {
	Type  CardType `json:"type"`
	Color Color    `json:"color"`
	Num   int8     `json:"numberValue"`
}

var allCards = map[Card]int8{}

func GenerateAllCards() {
	for color := Red; color <= Yellow; color++ {
		for i := 0; i < 10; i++ {
			if i != 0 {
				allCards[Card{Color: color, Num: int8(i), Type: Numeric}] = 2
				continue
			}
			allCards[Card{Color: color, Num: int8(i), Type: Numeric}] = 1
		}
		allCards[Card{Color: color, Type: Reverse, Num: -1}] = 2
		allCards[Card{Color: color, Type: Skip, Num: -1}] = 2
		allCards[Card{Color: color, Type: TakeTwo, Num: -1}] = 2
		allCards[Card{Color: color, Type: ChooseColor, Num: -1}]++
		allCards[Card{Color: color, Type: TakeFourChooseColor, Num: -1}]++
	}

	fmt.Println("All card added to map")
}

func GetOpponentHand(hand []Card, discard []Card) []Card {
	r := make(map[Card]int8, len(allCards))
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

func CanNextMoveFirst(c *Card, d *Card) bool {
	if c.Type == TakeFourChooseColor || c.Type == ChooseColor {
		return true
	}

	if c.Color != d.Color && c.Num != d.Num {
		// different color and number
		return false
	}

	return true
}

func CanNextMove(c *Card, d *Card) bool {
	if d.Type == Skip && c.Type != Skip {
		// skip on table and our card is not skip
		return false
	}

	if c.Type == ChooseColor || c.Type == TakeFourChooseColor {
		// black active
		return true
	}
	if c.Color != d.Color && c.Num != d.Num {
		// different color and number
		return false
	}
	// TODO: handle reverse
	if c.Type == Reverse || d.Type == Reverse {
		// active reverse
		return false
	}

	if c.Type == Skip && d.Type == Skip {
		// active skip
		return true
	}
	if c.Type == TakeTwo && (d.Type == TakeTwo || d.Type == TakeFourChooseColor) {
		// active take two
		return true
	}

	if d.Type == TakeTwo || d.Type == TakeFourChooseColor {
		// no active cards or cards with same number or color
		return false
	}

	return true
}

func (c *Card) String() string {
	return fmt.Sprintf("Color = %s, Type = %s, Number = %d", c.color(), c.cType(), c.Num)
}

func (c *Card) color() string {
	switch c.Color {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	case Yellow:
		return "Yellow"
	default:
		return "Black"
	}
}

func (c *Card) cType() string {
	switch c.Type {
	case Numeric:
		return "Numeric"
	case Reverse:
		return "Reverse"
	case Skip:
		return "Skip"
	case TakeTwo:
		return "TakeTwo"
	case TakeFourChooseColor:
		return "TakeFourChooseColor"
	case ChooseColor:
		return "ChooseColor"
	default:
		return "????"
	}
}
