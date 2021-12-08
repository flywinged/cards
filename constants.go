// Copyright Clayton Brown 2021. See LICENSE file.

package cards

// Values for determining other actions for the cards repository.
//	Users can change these values if they want.=
var N_SUITS = 4
var SUIT_NAMES = []string{"Clubs", "Diamonds", "Hearts", "Spades"}

var N_VALUES = 13
var VALUE_NAMES = []string{"Jack", "Queen", "King", "Ace"}

// Whether or not the highest value card in the deck doubles as the lowest value
// 	card in the deck.
var VALUE_WRAP = true

// Functions for change the suit and value names

// Adjust how many suits there are. When doing this, you must provide the names
// 	of those suits as well.
func AdjustSuits(suitNames []string) {
	N_SUITS = len(suitNames)
	SUIT_NAMES = suitNames
}