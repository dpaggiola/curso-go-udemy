package main

import (
	"fmt"
	"log"
	"github.com/dpaggiola/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Daniel", "Root", "Test"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	

	// message, err := greetings.Hello("Daniel")
	// if err != nil {
	//	log.Fatal(err)
	// }
	fmt.Println(messages)
}
