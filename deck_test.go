package deck

import (
	"testing"
)

func TestDeck_Init(t *testing.T) {
	var d Deck

	if 0 != d.Len() {
		t.Fatalf("deck should has 0 cards before Init, have: %d \n", d.Len())
	}

	d.Init()

	if 52 != d.Len() {
		t.Fatalf("deck should has 52 cards after Init, have: %d \n", d.Len())
	}
}

func TestDeck_Shuffle(t *testing.T) {
	for i := 0; i < 4; i++ {
		var d Deck
		var d2 Deck

		d.Init()
		d2.Init()
		d2.Shuffle()
		cards, _ := d.Draw(1)
		cards2, _ := d2.Draw(1)

		for i, card := range cards {
			if card.Equal(cards2[i]) {
				t.Fatalf("deck should be properly shuffled, index %d are equal", i)
			}
		}
	}
}

func TestDeck_Draw(t *testing.T) {
	var d Deck
	var prev []Card
	var cards []Card
	var err error

	d.Init()

	if _, err := d.Draw(53); err != OutOfLength {
		t.Fatalf("over draw should return an error OutOfLength, have: %v \n", err)
	}

	for i := 0; i < 52; i++ {
		cards, err = d.Draw(1)
		if err != nil {
			t.Fatal(err)
		}

		for i, card := range prev {
			if card.Equal(cards[i]) {
				t.Fatalf("draw should return new card each time, have: %v \n", card)
			}
		}

		prev = cards
	}

	if _, err := d.Draw(1); err != OutOfCards {
		t.Fatalf("over draw should return an error OutOfCards, have: %v \n", err)
	}
}

func TestDeck_Sharp(t *testing.T) {
	var d Deck
	var card *Card

	d.Init()

	card = d.Sharp("", 0)
	if card == nil {
		t.Fatal("deck should sharp any card")
	}

	card = d.Sharp(Clubs, Values[Ace])
	if card == nil {
		t.Fatal("deck should sharp valid card")
	}

	card = d.Sharp(Clubs, Values[Ace])
	if card != nil {
		t.Fatal("deck should not be able sharp twice")
	}

	if 50 != d.Len() {
		t.Fatalf("deck should has 50 cards, have: %d \n", d.Len())
	}
}

func TestDeck_Append(t *testing.T) {
	var d Deck
	d.Init()

	d.Append([]Card{
		{
			Suit:   Clubs,
			Symbol: Symbols[Clubs],
			Name:   Ace,
			Value:  Values[Ace],
		},
		{
			Suit:   Clubs,
			Symbol: Symbols[Clubs],
			Name:   Ace,
			Value:  Values[Ace],
		},
	})

	if 54 != d.Len() {
		t.Fatalf("deck should has 54 cards, have: %d \n", d.Len())
	}
}

func TestCard_IsFaceCard(t *testing.T) {
	for _, rank := range Ranks[:9] {
		card := Card{Name: rank.Name, Value: rank.Value}
		if card.IsFaceCard() {
			t.Fatalf("card with value %q name %q should not be facecard \n", card.Value, card.Name)
		}
	}

	for _, rank := range Ranks[9:12] {
		card := Card{Name: rank.Name, Value: rank.Value}
		if !card.IsFaceCard() {
			t.Fatalf("card with value %q name %q should be facecard \n", card.Value, card.Name)
		}
	}
}

func TestCard_IsAce(t *testing.T) {
	for _, rank := range Ranks[:12] {
		card := Card{Name: rank.Name, Value: rank.Value}
		if card.IsAce() {
			t.Fatalf("card with value %q name %q should not be an Ace \n", card.Value, card.Name)
		}
	}

	for _, rank := range Ranks[12:] {
		card := Card{Name: rank.Name, Value: rank.Value}
		if !card.IsAce() {
			t.Fatalf("card with value %q name %q should be an Ace \n", card.Value, card.Name)
		}
	}
}

func TestCard_Equal(t *testing.T) {
	card := Card{
		Suit:   Clubs,
		Symbol: Symbols[Clubs],
		Name:   Three,
		Value:  Values[Three],
	}
	card2 := Card{
		Suit:   Clubs,
		Symbol: Symbols[Clubs],
		Name:   Three,
		Value:  Values[Three],
	}
	card3 := Card{
		Suit:   Diamonds,
		Symbol: Symbols[Diamonds],
		Name:   Three,
		Value:  Values[Three],
	}
	card4 := Card{
		Suit:   Clubs,
		Symbol: Symbols[Clubs],
		Name:   Five,
		Value:  Values[Five],
	}

	if !card.Equal(card2) {
		t.Fatal("same cards should match", card, card2)
	}

	if card.Equal(card3) {
		t.Fatal("diff cards should not match", card, card3)
	}

	if card.Equal(card4) {
		t.Fatal("diff cards should not match", card, card4)
	}
}
