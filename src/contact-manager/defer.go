package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	defer file.Close()

	_, err = file.Write([]byte("Hola, Daniel"))
	if err != nil {
		fmt.Println(err)
		return
	}

}
