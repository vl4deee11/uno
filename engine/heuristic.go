package engine

import "math"

func heuristicsEstimation(node *Node) float32 {
	if len(node.Hand) == 0 {
		return math.MaxFloat32
	}

	topCard := node.TopFromTable()
	sameColor:=0
	for i:=0;i<len(node.Hand);i++{
		if node.Hand[i].Color == topCard.Color {sameColor++}
	}

	sameNumber := 0
	if topCard.Num > -1 {
		for i := 0; i < len(node.Hand); i++ {
			if node.Hand[i].Num == topCard.Num {
				sameNumber++
			}
		}
	}

	uselessCards := len(node.Hand) - sameNumber - sameColor
	goodCards := float32(sameColor+sameNumber)*0.5
	badCards := 0.35*float32(uselessCards)
	return goodCards - badCards
}