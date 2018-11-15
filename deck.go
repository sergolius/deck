package deck

import (
	"errors"
	"math/rand"
	"time"
)

var (
	OutOfCards  = errors.New("deck is out of cards")
	OutOfLength = errors.New("deck doesn't have enough cards")
)

const (
	Clubs    = "Clubs"
	Diamonds = "Diamonds"
	Hearts   = "Hearts"
	Spades   = "Spades"
)

var Symbols = map[string]string{
	Clubs:    "♣",
	Diamonds: "♦",
	Hearts:   "♥",
	Spades:   "♠",
}

const (
	Two   = "Two"
	Three = "Three"
	Four  = "Four"
	Five  = "Five"
	Six   = "Six"
	Seven = "Seven"
	Eight = "Eight"
	Nine  = "Nine"
	Ten   = "Ten"
	Jack  = "Jack"
	Queen = "Queen"
	King  = "King"
	Ace   = "Ace"
)

var Values = map[string]int{
	Two:   2,
	Three: 3,
	Four:  4,
	Five:  5,
	Six:   6,
	Seven: 7,
	Eight: 8,
	Nine:  9,
	Ten:   10,
	Jack:  11,
	Queen: 12,
	King:  13,
	Ace:   14,
}

var Suits = []Suit{
	{Clubs, Symbols[Clubs]},
	{Diamonds, Symbols[Diamonds]},
	{Hearts, Symbols[Hearts]},
	{Spades, Symbols[Spades]},
}

var Ranks = []Rank{
	{Two, Values[Two]},
	{Three, Values[Three]},
	{Four, Values[Four]},
	{Five, Values[Five]},
	{Six, Values[Six]},
	{Seven, Values[Seven]},
	{Eight, Values[Eight]},
	{Nine, Values[Nine]},
	{Ten, Values[Ten]},
	{Jack, Values[Jack]},
	{Queen, Values[Queen]},
	{King, Values[King]},
	{Ace, Values[Ace]},
}

type Deck struct {
	Cards []Card
}

type Card struct {
	Suit   string
	Symbol string
	Name   string
	Value  int
}

type Suit struct {
	Name   string
	Symbol string
}

type Rank struct {
	Name  string
	Value int
}

// Init initialize a deck
func (d *Deck) Init() {
	var cards []Card
	for _, suit := range Suits {
		for _, rank := range Ranks {
			cards = append(cards, Card{
				Suit:   suit.Name,
				Symbol: suit.Symbol,
				Name:   rank.Name,
				Value:  rank.Value,
			})
		}
	}
	d.Cards = cards
}

// Shuffle shuffle the deck
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// Draw draws `n` cards of deck
func (d *Deck) Draw(n int) ([]Card, error) {
	if d.Len() == 0 {
		return nil, OutOfCards
	}
	if d.Len() < n {
		return nil, OutOfLength
	}
	cards := d.Cards[:n]
	d.Cards = d.Cards[n:]
	return cards, nil
}

// Sharp card sharp
func (d *Deck) Sharp(suit string, value int) *Card {
	for i, card := range d.Cards {
		if (suit == "" || card.Suit == suit) && (value == 0 || card.Value == value) {
			d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
			return &card
		}
	}
	return nil
}

// Append adds cards to deck
func (d *Deck) Append(cards []Card) {
	d.Cards = append(d.Cards, cards...)
}

// Len returns len of deck
func (d Deck) Len() int {
	return len(d.Cards)
}

// IsAce
func (c Card) IsAce() bool {
	return 14 == c.Value
}

// IsFaceCard jacks, queens, kings
func (c Card) IsFaceCard() bool {
	return 10 < c.Value && c.Value < 14
}

// IsFaceCard jacks, queens, kings
func (c Card) Equal(card Card) bool {
	return c.Value == card.Value && c.Name == card.Name && c.Suit == card.Suit
}
