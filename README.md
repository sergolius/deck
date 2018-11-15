# Deck

Card deck

#### Initialize
```
var d Deck
d.Init()
```
#### Shuffle/Reshuffle
```
d.Shuffle()
```
#### Draw card number
```
card, err := d.Draw(4)
```
#### Sharp a card
```
card := d.Sharp(Clubs, Values[Ace])
card = d.Sharp(Diamonds, 0)
card = d.Sharp("", Values[Ace])
```