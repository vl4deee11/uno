package engine

type GameStatus uint8

const (
	InProcess GameStatus = iota
	Win
	Lose
)

type Color uint8

const (
	Red Color = iota
	Green
	Blue
	Yellow
)

type CardType uint8

const (
	Numeric CardType = iota
	Reverse
	Skip
	TakeTwo
	TakeFourChooseColor
	ChooseColor
)
