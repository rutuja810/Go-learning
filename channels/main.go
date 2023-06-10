package main

// Channels are typed i.e whatever we send through channel must be always of same type. They are used for communication between go routines
// We never share variables in go routine

// Go routine - separate line of code execution that can only handle blcoking code

// Goroutines are working under the main Goroutines if the main Goroutine terminated, then all the goroutine present in the program also terminated.
// Goroutine always works in the background.
import (
	"fmt"
	"net/http"
	"time"
)

func main()  {

	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://youtube.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
}

	c := make(chan string)

	for _,link := range links {
		go checkstatus(link, c)

	}

	 // Infinite for loop

	// for {                              
	// 	go checkstatus(<-c, c)
	// }
	
	// Alternative to above for loop to understand easily by new engineer 

	for l := range c {
		go func(link string) {                    //Function literal
			time.Sleep(5 * time.Second)
			checkstatus(link,c)
	}(l)       //pass l of for loop to the function literal
	
}
}

func checkstatus(link string, c chan string) {
	_, err := http.Get(link)
	//The identifier represented by the underscore character(_) is known as a blank identifier. It is used as an anonymous placeholder instead of a regular identifier
	if err != nil {
		fmt.Println(link, "is down")
		c <- link
		return
	}
	fmt.Println(link, "is up!!")
	c <- link
}