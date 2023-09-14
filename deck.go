package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type card struct {
	suit  string
	value string
}

// create a new type deck which is a slice of card struct
type deck []card

// receiver function of all deck type values
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card.toString())
	}
}

func newDeck() deck {
	newDeck := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, cardSuite := range cardSuites {
		for _, cardValue := range cardValues {
			newCard := card{suit: cardSuite, value: cardValue}
			newDeck = append(newDeck, newCard)
		}
	}
	return newDeck
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (c card) toString() string {
	return c.value + " of " + c.suit
}

func (d deck) toString() string {
	var allCardSlice []string
	for _, cardInDeck := range d {
		allCardSlice = append(allCardSlice, cardInDeck.toString())
	}
	return strings.Join([]string(allCardSlice), ",")
}

func (d deck) writeToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// get a []byte and error
	bs, err := os.ReadFile(filename)
	if err != nil {
		// log the error and quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	splitStringSplice := strings.Split(string(bs), ",")
	var cardSplice []card
	for _, splitString := range splitStringSplice {
		foundSuite, foundValue, _ := strings.Cut(splitString, " of ")
		cardSplice = append(cardSplice, card{suit: foundSuite, value: foundValue})
	}
	return deck(cardSplice)
}

func (d deck) shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
