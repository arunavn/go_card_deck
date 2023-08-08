package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)
	cards := []string{newCard(), newCard()}
	cards = append(cards, "Six of Spades")
	fmt.Println(cards)
	for _, x := range cards {
		fmt.Println(x)
	}

}

func newCard() string {
	return "Five of Diamond"
}
