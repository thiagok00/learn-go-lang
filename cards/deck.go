package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cardSuit string

// Represent a suit of a card
const (
	Spades   = "Spades"
	Diamonds = "Diamonds"
	Hearts   = "Hearts"
	Clubs    = "Clubs"
)

type cardValue string

// Represent a card value
const (
	Ace   = "Ace"
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
)

type card struct {
	cardSuit
	cardValue
}

type deck []card

func newDeck() deck {

	cardSuits := []cardSuit{Spades, Diamonds, Hearts, Clubs}
	cardValues := []cardValue{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

	d := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d = append(d, card{cardSuit: suit, cardValue: value})
		}
	}
	return d
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// func (d deck) toString() string {
// 	str := ""
// 	for _, card := range d {
// 		str = fmt.Sprintln()
// 	}
// 	// return strings.Join(d, ",")
// }

// func (d deck) saveToFile(filename string) error {
// 	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
// }

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// func newDeckFromFile(filename string) deck {
// 	bs, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		fmt.Println("Error:", err, bs)
// 		os.Exit(1)
// 	}
// 	return deck(strings.Split(string(bs), ","))
// }
