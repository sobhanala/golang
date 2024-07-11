package main

import "fmt"

func main() {

	messages := make(chan string)
	messages2 := make(chan string)
	go func() { messages <- "ping" }()
	go func() { messages2 <- "hi" }()

	msg := <-messages
	fmt.Println(msg)
}
