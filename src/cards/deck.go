package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

//Create a new type of deck
//which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) countCards() int {
	return len(d)
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d),",")
}

func (d deck) saveToFile(filename string) error{	
	return ioutil.WriteFile(filename, []byte(d.toString()), 0655)
}

func newDeckFromFile(filename string) (deck){
	byteslice, err := (ioutil.ReadFile(filename))
	if err != nil{
		fmt.Println("Error:",err)
		os.Exit(1) //this indicates that there were an error in our program
	}

	s := strings.Split(string(byteslice),",")//split a string
	return deck(s) 
		
}

func (d deck) shuffle(){
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	
	for i, _ := range d{
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition],d[i]
	}
}