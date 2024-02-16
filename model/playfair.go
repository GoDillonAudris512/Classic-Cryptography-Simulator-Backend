package model

type Bigram struct {
	First  rune
	Second rune
}

type Pair struct {
	Row int
	Col int
}

type MapPair struct {
	CharToPair map[rune]Pair
	PairToChar map[Pair]rune
}
