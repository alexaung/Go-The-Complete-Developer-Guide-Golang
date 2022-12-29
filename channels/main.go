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
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://codedezk.com",
		"http://thitsarparami.org",
		"http://myanmargita.com",
		"http://medium.com/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	// 1
	// for {
	// 	//fmt.Println(<-c)
	// 	go checkLink(<-c, c)
	// }
	// 2
	// for l := range c{
		
	// 	go checkLink(l, c)
	// }
	// 3
	for l := range c {
		go func (link string)  {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		//c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// c <- "Yep its up"
	c <- link
}
