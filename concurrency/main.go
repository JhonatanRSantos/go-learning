package main

import (
	"fmt"
	"net/http"
	"time"
)

type Information struct {
	Data interface{}
}

func main() {
	links := []string{
		"http://google.com.br",
		"http://facebook.com.br",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com.br",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-channel)
	// }

	// for {
	// 	go checkLink(<-channel, channel)
	// }

	// for c := range channel {
	// 	go checkLink(c, channel)
	// }

	for l := range channel {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, channel)
		}(l)
	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is down!")
		// channel <- link + " is down!"
	} else {
		fmt.Println(link, "is up!")
		// channel <- link + " is up!"
	}
	channel <- link
}
