package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://dummytextwontwrokatll.com",
		"http://golang.org",
		"http://instagram.com",
	}
	for _, link := range links {
		check(link)
	}
}

func check(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "isnt working")
		return
	}
	fmt.Println(url, "is up")
}
