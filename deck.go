package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of deck, of type slice of str
type deck []string

func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four"}
	for _, suit := range cardsSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, x := range d {
		fmt.Println(x)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {
	deckString := ""
	// for _, x := range d {
	// 	deckString = deckString + x + ","
	// }
	// return deckString[:len(deckString)-1]
	deckString = strings.Join(d, ",")
	return deckString
}

func (d deck) saveToFile(filePath string) {
	byteDeck := []byte(d.toString())
	os.WriteFile(filePath, byteDeck, 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffleDeck() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i, _ := range d {
		ri := r.Intn(len(d) - 1)
		d[i], d[ri] = d[ri], d[i]
	}
}
