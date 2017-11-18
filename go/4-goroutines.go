package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go my_slow_name("John Doe")
	fmt.Println("Name finished\n\n")
	var wait string
	fmt.Scanln(&wait)
}

func my_slow_name(name string) {
	letters := strings.Split(name, "")

	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}
