package main

import (
	"fmt"
	"github.com/dpaggiola/greetings"
)

func main() {
	message := greetings.Hello("Daniel")
	fmt.Println(message)
}