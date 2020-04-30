package main

import "fmt"

func main() {

	// CHECK:1 How to declare a variable, there are two ways

	// Way - 1
	// variable a of type int
	var a int
	a = 20

	// Way - 2
	b := "Hello i'm second way of declaring variables"

	c := newCard()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// CHECK: 2 How to declare arrays and how to loop over them

	cards := []string{"Hey", newCard(), newCard()}

	fmt.Println("Array", cards)

	// Loop over a array or slice

	for i, card := range cards {
		fmt.Println(i, card)
	}

	// How to add an element to Array
	cards = append(cards, "Hey im appending a new card")

	for i, c := range cards {
		fmt.Println(i, c)
	}

}

// How to declare a funtion and use

// Always return type of a function

func newCard() string {
	return "Hey new card"
}
