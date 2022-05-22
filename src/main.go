package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Hello World")

	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Declaracion de variables enteras
	base := 12 // Para crear de una vez lleva los 2 puntos
	var altura int = 14
	var area int

	fmt.Println("base", base)
	fmt.Println("altura", altura)
	fmt.Println("area", area)

	// Zero values -- go no asigna null ni undefined asigna valores por default
	var a int     // 0
	var b float64 // 0
	var c string  // vacio
	var d bool    // false

	fmt.Println(a, b, c, d)

	// Ejercicio
	// Calcular area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("areaCuadrado", areaCuadrado)

	// Operadores aritmeticos
	x := 10
	y := 50

	result := x + y
	fmt.Println("suma", result)

	result = y - x
	fmt.Println("resta", result)

	result = y * x
	fmt.Println("multiplicacion", result)

	result = y / x
	fmt.Println("divicion", result)

	result = y % x
	fmt.Println("modulo o residuo", result)

	x++
	fmt.Println("inclemental", x)

	x--
	fmt.Println("decremental", x)

	// area del rectangulo
	base2 := 4
	altura2 := 2
	println("Area del rectangulo ", base2*altura2)

	// area del trapecio
	base3 := 3
	baseAbajo := 2
	areaTrapecio := altura * (base3 + baseAbajo) / 2
	println("Area del trapecio ", areaTrapecio)

	// area del circulo
	radio := 2
	areaCirculo := math.Pi * math.Pow(float64(radio), 2)
	println("Area del circulo", areaCirculo)

	// Tipos de datos primitivos
	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^16-1
	//uint32 = 32bits = 0 a 2^32-1
	//uint64 = 64bits = 0 a 2^64-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

	helloMessage := "Hello"
	worldMessage := "world"

	// Println: Salto de Linea Automatico
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	// Con valores seguros
	// %s hace referencia a un string
	// %d hace referencia a un numero
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	// Con valores inseguros
	// %v Si no se sabe el tipo de dato
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Saber el tipo de datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
	fmt.Printf("2: %T", 2)
}
