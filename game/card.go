package game

type CardColor int
const (
	Red CardColor = iota
	Blue
	Green
	Yellow
	Wild
)

type CardType int
const (
	Number CardType = iota
	Skip
	Reverse
	DrawTwo
	WildCard
	WildDrawFour
)

type Card struct{
	Color CardColor
	Type CardType
	Value int // Only used for number cards (0-9)
}
