package main

import (
	"fmt"
	"math"

	"rsc.io/quote"
)

// Declaración de constantes
const Pi = 3.14
const (
	x = 100
	y = 0b1010 // binario
	z = 0o12 // octal
	w = 0xFF // hexadecimal
)

const (
	Domingo = iota + 1 // iota inicia desde cero  e incrementa en 1 en cada constante
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Hello())

	fmt.Println(math.MinInt64, math.MaxInt64)

	fullName := "Alex Roel \t(alias \"roelcode\")\n"
	fmt.Println(fullName)

	var a byte = 'a'
	fmt.Println(a)

	s := "hola"
	fmt.Println(s[0])

	var r rune = '♡'

	fmt.Println(r)

	var (
		defaultInt int
		defaultUint uint
		defaultFloat float32
		defaultBool bool
		defaultString string
	)

	fmt.Println(defaultInt, defaultUint, defaultFloat, defaultBool, defaultString)

	// Declaración de variables
	// var firstName, lastName string
	// var age int

	// var (
	// 	firstName, lastName string
	// 	age                 int
	// )

	// firstName = "Alex"
	// lastName = "Roel"
	// age = 27

	// Usando var, se declaran las variables de forma global
	// var firstName, lastName, age = "Alex", "Roel", 27

	// Usando :=, se declaran las variables de forma local
	firstName, lastName := "Alex", "Roel"
	age := 27

	fmt.Println(firstName, lastName, age)
	fmt.Println(Pi)

	fmt.Println(x, y, z, w)
	fmt.Println(Viernes)

}
