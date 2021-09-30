package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println("error")
		return
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}
