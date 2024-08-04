package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error is :", err)
		os.Exit(1)
	}
	bs := make([]byte, 40000)
	resp.Body.Read(bs)
	s := string(bs)
	os.Remove("responseBody")
	os.WriteFile("responseBody", []byte(s), 0666)
}
