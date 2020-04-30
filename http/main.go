package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// Lets see how to make http calls
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	fmt.Println(resp)

	// To get the response from the request, we need to do something
	// Create a byte slice of hevay capacity just to make sure that
	// all resposne fit in
	// bs := make([]byte, 99999)

	// Now pass this bs to body resp, where all the content will be filled
	// in this bs
	// just printing bs into a single string
	// resp.Body.Read(string(bs))

}
