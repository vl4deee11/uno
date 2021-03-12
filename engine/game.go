package engine

import (
	"math"
	"math/rand"
)

type Node struct {
	Hand         []Card
	OpponentHand []Card
	Curr         *Card
	Nodes        []*Node
	Parent       *Node
	Lvl          uint8
	MinE         float64
	MaxE         float64
}

func (n *Node) Add(c *Node) {
	n.Nodes = append(n.Nodes, c)
}

func (n *Node) AddE(e float64)  {
	n.MinE = math.Min(n.MinE, e)
	n.MaxE = math.Max(n.MaxE, e)
}


func GetNextCard(hand []Card, opponent []Card, curr *Card) *Card {
	root := &Node{
		Hand:         hand,
		OpponentHand: opponent,
		Curr:         curr,
		Nodes:        make([]*Node, 0, len(hand)),
		Lvl:          0,
		Parent:       nil,
	}

	s := make([]*Node, 0)
	for i := 0; i < len(hand); i++ {
		if CanNextMove(&hand[i], curr) {
			n:=&Node{
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

		for i := 0; i < len(l); i++ {
			if CanNextMove(&l[i], node.Curr) {
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
				s = append(s, next)
			}
		}
		if isLeaf {
			leafs = append(leafs, node)
		}
	}

	estimations := map[float64]*Node{}
	for len(leafs) > 0 {
		node := leafs[len(leafs)-1]
		leafs[len(leafs)-1] = nil
		leafs = leafs[:len(leafs)-1]
		if node.Parent == nil {
			continue
		}
		if len(node.Nodes) == 0 {
			e := heuristicsEstimation(node)
			if node.Lvl&1 == 0 {
				e = -e
			}
			estimations[e] = node
			node.Parent.AddE(e)
		} else if node.Lvl&1 == 0 {
			estimations[node.MinE] = node
			node.Parent.AddE(node.MinE)
		} else {
			estimations[node.MaxE] = node
			node.Parent.AddE(node.MaxE)
		}
		leafs = append(leafs, node.Parent)
	}

	if _, ok := estimations[root.MaxE]; !ok {
		var max float64 = 0
		var bestCard *Card
		for i := 0; i<len(root.Nodes);i++{
			if root.Nodes[i].MaxE > max {
				max = root.Nodes[i].MaxE
				bestCard = root.Nodes[i].Curr
			}
		}
		return bestCard
	}
	return estimations[root.MaxE].Curr
}

func heuristicsEstimation(node *Node) float64 {
	return rand.Float64()*700
}

func remove(s []Card, i int) []Card {
	f := make([]Card, len(s))
	copy(f, s)
	f[i] = f[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return f[:len(s)-1]
}
