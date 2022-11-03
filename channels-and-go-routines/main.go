package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://crt.sh",
		"https://stackoverflow.com",
		"https://github.com",
	}

	checkStatus(links)
}

func checkStatus(ln []string) {
	for i, l := range ln {
		_, err := http.Get(l)

		if err != nil {
			fmt.Println(i, l, "might be done")
		}

		fmt.Println(i, l, "is up")
	}
}
