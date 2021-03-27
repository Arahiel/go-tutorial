package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return strings.Split(string(bs), ",")
}

func (d deck) shuffle() {
	for i := range d {
		seed := time.Now().UnixNano()
		randomIndex := rand.New(rand.NewSource(seed)).Intn(len(d))
		d.swap(i, randomIndex)
	}
}

func (d deck) swap(first, second int) error {
	if first >= len(d) || second >= len(d) {
		return fmt.Errorf("ArgumentOutOfRangeException: Index was out of range. Must be non-negative and less than the size of the collection")
	}

	d[first], d[second] = d[second], d[first]
	return nil
}
