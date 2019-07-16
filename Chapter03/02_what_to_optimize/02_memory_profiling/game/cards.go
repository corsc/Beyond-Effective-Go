package game

import (
	"math/rand"
	"time"
)

// Shuffle randomizes the order of the supplied cards
func Shuffle(cards []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(a, b int) {
		cards[a], cards[b] = cards[b], cards[a]
	})

	waste := 0
	for x := 0; x < 1000000; x++ {
		// waste some CPU
		waste++
	}
	println(waste)
}

// Card represents a single playing card
type Card struct {
	Suit string
	Face string
}

// NewDeck creates a deck of cards (NOTE: the cards are in predictable order)
func NewDeck() []Card {
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
