package engine

type Node struct {
	hand         []*Card
	opponentHand []*Card
	curr         *Card
	nodes        []*Node
}
