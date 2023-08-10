package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)
	// cards = append(cards, "Six of Spades")
	// cards := newDeck()
	cards := newDeckFromFile("test.txt")
	// cards.saveToFile("test.txt")

	// for _, x := range cards {
	// 	fmt.Println(x)
	// }
	cards.print()
	cards.shuffleDeck()
	fmt.Println("Hand................................")
	cards.print()
	// fmt.Println("Hand................................")
	// hand, remainingCards := deal(cards, 4)
	// hand.print()
	// fmt.Println("Remaining...........................")
	// remainingCards.print()
}

func newCard() string {
	return "Five of Diamond"
}
