package main

import (
	"fmt"
)

func test_channel(my_channel chan string) {

	for {
		var name string
		fmt.Printf("\nInsert message: ")
		fmt.Scanln(&name)
		my_channel <- name
	}

}

func main() {
	fmt.Printf("\n===========================\n")

	my_channel := make(chan string)

	go test_channel(my_channel)

	for {
		message := "\nYour message is( " + <-my_channel + " )\n"
		fmt.Println(message)
	}

	fmt.Printf("\n===========================\n")
}
