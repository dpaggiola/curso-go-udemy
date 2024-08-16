package main

import (
	"fmt"
	"math"
	"strconv"

	"rsc.io/quote"
)

// Declaración de constantes
const Pi = 3.14
const (
	x = 100
	y = 0b1010 // binario
	z = 0o12   // octal
	w = 0xFF   // hexadecimal
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
	fmt.Println(quote.Go())

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
		defaultInt    int
		defaultUint   uint
		defaultFloat  float32
		defaultBool   bool
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

	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(int32(integer16) + integer32)

	st := "100"
	iv, _ := strconv.Atoi(st)

	fmt.Println(iv + 1)

	n := 42
	st = strconv.Itoa(n)
	fmt.Println(st + st)

	name1 := "Alex"
	age1 := 28

	fmt.Printf("Hola, me llamo %s y tengo %d años. \n", name1, age1)

	greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años. \n", name1, age1)

	fmt.Println(greeting)

	var name2 string
	var age2 int

	fmt.Println("Ingrese su nombre: ")
	fmt.Scanln(&name2)
	fmt.Println("Ingrese su edad: ")
	fmt.Scanln(&age2)

	fmt.Printf("Hola, te llamas %s y tienes %d años. \n", name2, age2)
}
