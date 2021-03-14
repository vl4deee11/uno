package engine

import (
	"math/rand"
)

type Node struct {
	Hand         []Card
	OpponentHand []Card
	Curr         *Card
	Nodes        []*Node
	Parent       *Node
	Lvl          uint8
	estims       []*Pair
}

type Pair struct {
	E float32
	C *Card
}

func (n *Node) Add(c *Node) {
	n.Nodes = append(n.Nodes, c)
}

func (n *Node) AddE(e float32, c *Card) {
	n.estims = append(n.estims, &Pair{e, c})
}

func GetNextCard(hand []Card, opponent []Card, curr *Card) []Card {
	root := &Node{
		Hand:         hand,
		OpponentHand: opponent,
		Curr:         curr,
		Nodes:        make([]*Node, 0, len(hand)),
		Lvl:          0,
		Parent:       nil,
	}

	s := make([]*Node, 0)
	usedCards := map[Card]struct{}{}
	for i := 0; i < len(hand); i++ {
		if _, ok := usedCards[hand[i]]; !ok && CanNextMove(&hand[i], curr) {
			n := &Node{
				Hand:         remove(hand, i),
				Curr:         &hand[i],
				OpponentHand: opponent,
				Nodes:        make([]*Node, 0, len(hand)),
				Lvl:          1,
				Parent:       root,
			}
			root.Add(n)
			s = append(s, n)
		}
	}

	leafs := make([]*Node, 0)
	for len(s) > 0 {
		node := s[len(s)-1]
		s[len(s)-1] = nil
		s = s[:len(s)-1]
		if node.Lvl == 3 {
			leafs = append(leafs, node)
			continue
		}

		isLeaf := true
		var l []Card
		if node.Lvl == 1 {
			l = opponent
		} else {
			l = node.Parent.Hand
		}

		usedCards := map[Card]struct{}{}
		for i := 0; i < len(l); i++ {
			if _, ok := usedCards[l[i]]; !ok && CanNextMove(&l[i], node.Curr) {
				isLeaf = false
				next := &Node{
					Hand:         remove(l, i),
					Curr:         &l[i],
					Nodes:        make([]*Node, 0, len(l)),
					Lvl:          node.Lvl + 1,
					Parent:       node,
					OpponentHand: node.Hand,
				}
				node.Add(next)
				usedCards[l[i]] = struct{}{}
				s = append(s, next)
			}
		}
		if isLeaf {
			leafs = append(leafs, node)
		}
	}

	usedNodes := map[*Node]struct{}{}
	for len(leafs) > 0 {
		node := leafs[0]
		leafs[0] = nil
		leafs = leafs[1:]
		if _, ok := usedNodes[node]; ok {
			continue
		}
		usedNodes[node] = struct{}{}
		if node.Parent == nil {
			continue
		}
		if len(node.Nodes) == 0 {
			e := heuristicsEstimation(node)
			if node.Lvl&1 == 0 {
				// even
				e = -e
			}
			node.Parent.AddE(e, node.Curr)
		} else if node.Lvl&1 == 0 {
			// even
			node.Parent.AddE(Max(node.estims).E, node.Curr)
		} else {
			node.Parent.AddE(Min(node.estims).E, node.Curr)
		}

		leafs = append(leafs, node.Parent)
	}

	best := Max(root.estims)
	if best != nil {
		if best.C.Type != Numeric {
			return []Card{*best.C}
		}
		return groupCards(*best.C, hand)
	}
	return []Card{}
}

func heuristicsEstimation(node *Node) float32 {
	return rand.Float32() * 700
}

func groupCards(bestCard Card, hand []Card) []Card {
	var r []Card
	for i := 0; i < len(hand); i++ {
		if hand[i].Color == bestCard.Color && hand[i].Type == Numeric {
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
