package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://amazon.com",
	}

	//creating a channel
	c := make(chan string)

	for _, link := range links {
		go getResponse(link, c) // creating a new go routine
	}

	//wait for routines to complete
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func getResponse(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("Error:", err)
		c <- link + " is broken"
		return
	}

	c <- link + " is up"
}
