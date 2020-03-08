package main

import "fmt"

func main() {
	//var card string = "Ace of spades"
	// card := "Ace of Spades"
	// card = "King of Spades"

	// card2 := newCard()

	// cards := deck{"1 of Diamonds", "2 of Spades"}
	// cards = append(cards, "3 of Hearts")
	// fmt.Println(card)
	// fmt.Println(card2)

	cards := newDeck()
	hand, remainingCards := deal(cards, 4)
	hand.print()
	remainingCards.print()

	cards.saveToFile("allCards")
	cards = readDeckFromFile("allCards")
	fmt.Println("Printing cards from file")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
