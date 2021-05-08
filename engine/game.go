package engine

import (
	"math"
)

type Node struct {
	Hand         []Card
	OpponentHand []Card
	table        []*Card
	Parent       *Node
	Lvl          uint8
}

func (n *Node) PutOnTable(c *Card) {
	n.table = append(n.table, c)
}

func (n *Node) TopFromTable() *Card {
	return n.table[len(n.table)-1]
}

// min score for my turn
var alpha float32 = 0

// max score for opponent
var beta float32 = 0

func GetNextCard(hand []Card, opponent []Card, curr *Card) []Card {
	alpha = -math.MaxFloat32
	beta = math.MaxFloat32

	usedCards := map[Card]struct{}{}

	var max float32 = -math.MaxFloat32
	maxCard := Card{Num: -11}

	for i := 0; i < len(hand); i++ {
		if _, ok := usedCards[hand[i]]; !ok && CanNextMoveFirst(&hand[i], curr) {
			usedCards[hand[i]] = struct{}{}
			next := &Node{
				Hand:         remove(hand, i),
				table:        []*Card{curr},
				OpponentHand: opponent,
				Lvl:          1,
				Parent:       nil,
			}
			next.PutOnTable(&hand[i])
			e := DFS(next, false)
			if e >= max {
				max = e
				maxCard = hand[i]
			}
		}
	}

	if maxCard.Num != -11 {
		if maxCard.Type == Numeric {
			groupedByColor := groupCardsByColor(maxCard, hand)
			groupedByNumber := groupCardsByNumber(maxCard, hand)
			if len(groupedByColor) > len(groupedByNumber) {
				return groupedByColor
			}
			return groupedByNumber
		}
		return []Card{maxCard}
	}

	return []Card{}
}

func DFS(node *Node, maximizationStep bool) float32 {
	if node.Lvl == 10 {
		return heuristicsEstimation(node)
	}
	var l []Card
	if node.Lvl == 1 {
		l = node.OpponentHand
	} else {
		l = node.Parent.Hand
	}

	var resEstimation float32 = math.MaxFloat32
	if maximizationStep {
		resEstimation = -resEstimation
	}

	usedCards := map[Card]struct{}{}
	for i := 0; i < len(l); i++ {
		if _, ok := usedCards[l[i]]; !ok && CanNextMove(&l[i], node.TopFromTable()) {
			usedCards[l[i]] = struct{}{}
			next := &Node{
				Hand:         remove(l, i),
				table:        node.table,
				Lvl:          node.Lvl + 1,
				Parent:       node,
				OpponentHand: node.Hand,
			}
			next.PutOnTable(&l[i])

			}
		}
	}

		}
		return e
	}

	return resEstimation
}

func groupCardsByColor(bestCard Card, hand []Card) []Card {
	var r []Card
	for i := 0; i < len(hand); i++ {
		if hand[i].Color == bestCard.Color && hand[i].Num == bestCard.Num && hand[i].Type == Numeric {
			r = append(r, hand[i])
		}
	}
	return r
}

func groupCardsByNumber(bestCard Card, hand []Card) []Card {
	var r []Card
	for i := 0; i < len(hand); i++ {
		if hand[i].Num == bestCard.Num && hand[i].Type == Numeric {
			r = append(r, hand[i])
		}
	}
	return r
}

func remove(s []Card, i int) []Card {
	f := make([]Card, len(s))
	copy(f, s)
	f[i] = f[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return f[:len(s)-1]
}
