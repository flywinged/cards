// Copyright Clayton Brown 2021. See LICENSE file.

package cards

// Card is a simple struct which contains information about
//	what value and suit a card has.
type Card struct {

	Value uint8
	Suit uint8

}

// Some helper methods for deplaying card data.

// Determines the suit of the card
func (c *Card) GetSuit() string {
	return SUIT_NAMES[c.Suit]
}

// Determines the value of a card
func (c *Card)

// Determines the name of the card
func (c *Card) GetString() string {

	

}