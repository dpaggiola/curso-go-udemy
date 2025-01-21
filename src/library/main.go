package main

import (
	"library/animal"
	// "library/book"
)

func main() {
	// myBook := book.NewBook("The Art of Computer Programming", "Donald Knuth", 700)
	
	// myBook.SetTitle("New Title")
	// fmt.Println(myBook.GetTitle())
	
	// myTextBook := book.NewTextBook("The Go Programming Language", "Alan A. A. Donovan & Brian W. Kernighan", "Addison-Wesley", "Intermediate", 380)
	
	// // myBook.PrintInfo()
	// // myTextBook.PrintInfo()
	// book.Print(myBook)
	// book.Print(myTextBook)

	// miPerro := animal.Perro{Nombre: "Firulais"}
	// miGato := animal.Gato{Nombre: "Michi"}

	// animal.HacerSonido(&miPerro)
	// animal.HacerSonido(&miGato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Firulais"},
		&animal.Gato{Nombre: "Michi"},
		&animal.Perro{Nombre: "Rex"},
		&animal.Gato{Nombre: "Garfield"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}