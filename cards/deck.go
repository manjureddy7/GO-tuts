package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Create a type of 'deck' which extends the functionality of slice of strings

type deck []string

// Create and return a list of playing cards. Essentially list of strings
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// The below function looks bit wierd
// It indicates any variable of type 'deck' can access print function
// d indiactes the variable which is calling the print function
// in main.go cards variable is calling print function
// so d refers to cards

// Prints the content of the passed deck
func (d deck) print(msg string) {
	for _, card := range d {
		fmt.Println(msg + " " + card)
	}
}

// Returns subset of deck (picked cards, remaining cards)

func deal(d deck, handSize int) (deck, deck) {

	// To get subset of a slice we have inbuilt methods as
	// slice[FromIndexIncluding:ToIndexExcluding]
	// slice[0:2] -> Give me elements from starting 0 to 2 but exclude 2nd index
	// fruits = {"banana","apple","orange","grapes"}
	// fruits[0:2] -> 0,1 which are banana & apple
	// fruits[:2] -> go infers this as from beginning to 2nd index , banana & apple
	// fruits[2:] -> from 2nd index to end -> orange, grapes

	return d[:handSize], d[handSize:]
}

// To save our decks into a file, we need to convert our deck into
// Byte slice, because the data in file should be saved in byte
// deck -> []string -> string -> []byte

// lets convert our deck to string
func (d deck) toString() string {
	// First convert our deck to []string -> []string(deck)
	// Convert to single string by Join(your sting slice, bywantyouwantseperate)
	stringSlice := []string(d)
	return strings.Join(stringSlice, ",")
}

// Now lets save our deck to file
func (d deck) saveToFile(fileName string) error {
	singleString := d.toString()
	// 0666 means a permission that anyone can read or write file
	return ioutil.WriteFile(fileName, []byte(singleString), 0666)
}

// Read file from harddrive/system

func newDeckFromFile(filename string) deck {

	resp, err := ioutil.ReadFile(filename)

	// The common check if the operation has error occured
	// If nothing went wrong the error will have value of NIL (nil)
	// Means there's an error, so handle it
	if err != nil {
		// Log the error
		fmt.Println("Error:", err)
		// To quit the program completely
		// os.Exit(1)
	}

	// The resposnse is in form of []byte
	// We need to convert it to deck []byte -> deck
	// []byte -> string -> []string -> deck

	// First byte[] to string
	// We can do this by type conversion
	// string(resp)
	returnedDeck := strings.Split(string(resp), ",")
	return deck(returnedDeck)

}

// Shuffle all the cards in the deck
func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		// Generate a number between 0, and length of slice
		// Swap the index of each current card with new position
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
