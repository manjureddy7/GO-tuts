package main

import "fmt"

// A type is called concrete type if it can returns value out of it
// ex: struct, map, slices

// An interface type is something which cannot returns value out of it

// define interface here
type bot interface {
	getGreeting() string

	// Basically this interface type bot is telling everyone in the
	// package that if any type has a method getGreeting() which returns
	// string, now you are a part of this interface
	// funnishly spekaing now you are a type of bot
	// you can access any function with type bot and excutes it

	// Lets say it also has one more method
	// getUsersCount() int

	// if our englishBot & spanishBot doesnt have getUsersCount
	// then those two are not part of BOT.

	// So any other type to be part of BOT they should contain all the
	// methods of interface type BOT
	// SHOULD & MUST contain methods
}

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	// lets learn INTERFACES here

	// Avatar is an english boy
	english := englishBot{}
	fmt.Println(english.getGreeting())
	fmt.Println("Hey Name method", english.getName())

	//Narcos is a spanish boy
	spanish := spanishBot{}
	fmt.Println(spanish.getGreeting())

	fmt.Println("english greet is", printGreet(english))
	fmt.Println("spanish greet is", printGreet(spanish))
}

// Consider the below getGreeting methods have very distinct way of implementation

// Lets create a custom English greeting generator
func (eb englishBot) getGreeting() string {
	// Consider this has very custom logic othern than spanish getGreeting
	return "Hey There!!"
}

// Lets create a custom Spanish greeting generator
func (sp spanishBot) getGreeting() string {
	// Consider this has very custom logic othern than English getGreeting
	return "Holaaa!"
}

// Lets craete new method for englishBot struct
// this specifies a struct can have many methods irrespective of what the
// Interface says, but the struct should contain the interface asked methods
func (eb englishBot) getName() string {
	return "Hey English Bot"
}

// In this scenario lets consider both printMeetings has identical implementation
// We can just reuse one function for both of the structs
// But we cant do that since receiver type of these two fucntions are diff.
// Interface tries to solve one of this problem

// func (p englishBot) printGreeting() string {
// 	return p.getGreeting()
// }

// func (p spanishBot) printGreeting() string {
// 	return p.getGreeting()
// }

// The above two statements can also be called as
// func printGreet(eb englishBot) {
// 	return eb.getGreeting()
// }

// func printGreet(sb spanishBot) {
// 	return sb.getGreeting()
// }

// As we observed above even though the implementation logic is same
// we are jst duplicating our code to suit different types of struct
// Lets see how interface handles it
// look in top of interface

func printGreet(b bot) string {
	return b.getGreeting()
}
