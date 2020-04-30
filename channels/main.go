package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Here we will list out some of links and make call to those links

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}
	// here we are looping over links and checking if the websites are up or down
	// here the function will take each link one at a time
	// & call that link (usually takes some time to get resp from that link)
	// once we got respose then head over to next link
	// Here we dont have a choice to make call to all the links AT ONCE
	// GO ROUTINE HELPS IN THAT CASE

	// To create a new Go routine add go infront of function call

	// Main go routine doesnt wait for the child routines to complete
	// their work, if main go routine compleets its work
	// it exits, wont wait at all
	// CHANNELS provide communication between all the routines
	// channel tells the main go routine that there are child go routines
	// you should look after

	// Make channel
	// This channel can communicate through type string
	// chan (string or int or float) can be anything
	// To send a value to channel
	// channel <- 5 , here sending a value of 5 into the channel
	// Receiving a value from channel
	// msgReceivedFromChannel <- channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	// Lets to continous pingning to each link
	// 	// so that it again creates new routine
	// 	link := <-c
	// 	go checkLink(link, c)
	// 	fmt.Println(<-c)
	// }

	// This is a kind of infinite loop
	// for {
	// 	// Lets to continous pingning to each link
	// 	// so that it again creates new routine
	// 	link := <-c
	// 	go checkLink(link, c)
	// 	// fmt.Println(<-c)
	// }

	// The above function can also be called as
	// Link will be returned from channel through all the routines

	for link := range c {
		// Now we want to pause our operation for the next fetch
		// for some period of time
		// we can do that  by
		// This is like anonymous fn in JS
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(link)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}
	fmt.Println(link, "is up & running!!!")
	c <- link
}
