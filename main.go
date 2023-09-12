package main

import "fmt"

func main() {
	// create a slice card deck
	fmt.Println("Creating new deck of 16 cards")
	cards := newDeck()
	fmt.Println("Printing current cards in deck")
	cards.print()
	fmt.Println("Shuffling cards in the deck...")
	cards.shuffle()
	fmt.Println("Dealing 5 cards")
	hand, remainingCards := deal(cards, 5)
	fmt.Println("Printing current cards in hand")
	hand.print()
	fmt.Println("Writing remaining cards to file my_cards...")
	remainingCards.writeToFile("my_cards")
	fmt.Println("Creating new deck from remaining cards...")
	newDeck := newDeckFromFile("my_cards")
	fmt.Println("Printing cards in new deck")
	newDeck.print()
	fmt.Println("Shuffling cards in new deck...")
	newDeck.shuffle()
	fmt.Println("Printing cards in shuffled deck")
	newDeck.print()
	fmt.Println("Dealing 5 cards for second hand")
	secondHand, remainingCards := deal(newDeck, 5)
	fmt.Println("Printing current cards in second hand")
	secondHand.print()
	fmt.Println("2 players READY TO PLAY")
}
