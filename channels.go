package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://bing.com",
		"http://golang.org",
	}
	/*
		for _, link := range links {
			checkLink(link)
		}
	*/
	//go routines : add go keywork infront of a function

	//Go scheduler runs one routine until it finishes or makes a blocking call
	//uses only 1 cpu core by default
	//with multiple cores, multiple routines can be simultaneously spun

	//as soon as main routine exits it doesnt wait for child routines to finish
	// we use channels to help with this
	// channels are used to communicate between go routines

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	//wait for a value to be sent into the channel "abc" <- channel
	//fmt.Println(<-c) //blocking line of code, and as soon as even one msg comes, this goes through

	//wait for these many strings coming into channel
	// for i := 0; i < len(links)-1; i++ {
	// 	fmt.Println(<-c)
	// }

	// for { //infinite loop
	// 	go checkLink(<-c, c)
	// }

	//alternate syntax
	for l := range c { //you can run for loop on range of channel like a slice
		//	time.Sleep(5 * time.Second) //will block main routine for 5sec
		//	go checkLink(l, c)
		go func(lnk string) {
			time.Sleep(5 * time.Second)
			checkLink(lnk, c)
		}(l) //pass by value
	}

}

func checkLink(link string, c chan string) {
	//time.Sleep(5 * time.Second) //will block go routine for 5sec
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link //sending string to the channel
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
