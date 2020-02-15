package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
}

// Output:
// Ace of Hearts

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Errorf("Wrong number of cards in the new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	wanted := Card{Rank: Ace, Suit: Spade}
	if cards[0] != wanted {
		t.Error("Expected Ace of Spades. Received: ", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	wanted := Card{Rank: Ace, Suit: Spade}
	if cards[0] != wanted {
		t.Error("Expected Ace of Spades. Received: ", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	shuffleRand := rand.New(rand.NewSource(0))
	orig := New()
	first := orig[40]
	second := orig[35]
	cards := New(Shuffle)
	if cards[0] != first {
		t.Error("Expected the first card to be %s, received %s", first, cards[0])
	}
	if cards[1] != second {
		t.Error("Expected the second card to be %s, received %s", second, cards[1])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, val := range cards {
		if val.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Wrong nuber of Joker suit cards")
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Error("Expected two's and three's to be filtered out")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 13*4*3 {
		t.Error("incomplete Deck of cards")
	}
}
