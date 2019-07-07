package shared

import (
	"math/rand"
	"net/http"
	"time"
)

func CardShuffler(resp http.ResponseWriter, req *http.Request) {
	// create a deck of cards
	cards := newDeck()

	// shuffle the cards
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(a, b int) {
		cards[a], cards[b] = cards[b], cards[a]
	})

	// return the result
	for index, card := range cards {
		if index > 0 {
			_, _ = resp.Write([]byte(", "))
		}
		_, _ = resp.Write([]byte(card.Face))
		_, _ = resp.Write([]byte(card.Suit))
	}
}

type Card struct {
	Suit string
	Face string
}

func newDeck() []Card {
	cards := make([]Card, 52)

	index := 0
	for _, suit := range []string{"♣", "♡", "♢", "♠"} {
		for _, face := range []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"} {
			cards[index] = Card{
				Suit: suit,
				Face: face,
			}
			index++
		}
	}

	return cards
}
