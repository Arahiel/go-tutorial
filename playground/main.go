package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	linkList := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
	}

	c := make(chan string, 1)

	for _, link := range linkList {
		go handleLink(link, c)
	}

	for {
		time.Sleep(time.Millisecond * 100)
	}
}

func handleLink(link string, c chan string) {
	checkLink(link, c)
	time.Sleep(time.Second * 5)
	go handleLink(<-c, c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
	}
	fmt.Println(link, " is OK!")
	c <- link
}
