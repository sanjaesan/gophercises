//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
)

//Suit ...
type Suit uint8

//Rank ...
type Rank uint8

var suits = [...]Suit{Spade, Diamond, Club, Heart}

//const ...
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

//const ...
const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

//const ...
const (
	minRank = Ace
	maxRank = King
)

//Card ...
type Card struct {
	Suit
	Rank
}

//New ...
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

//String ...
func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

//DefaultSort ...
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

//Less ...
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

//Sort ...
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

var shuffleRand = rand.New(rand.NewSource(0))

//Shuffle ...
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	perm := shuffleRand.Perm(len(cards))
	for i, val := range perm {
		ret[i] = cards[val]
	}
	return ret
}

//Jokers ...
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Suit: Joker,
				Rank: Rank(i),
			})
		}
		return cards
	}
}

//Filter ...
func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, card := range cards {
			if !f(card) {
				ret = append(ret, card)
			}
		}
		return ret
	}
}

//Deck ...
func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
