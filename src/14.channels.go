package main

import "fmt"

func say(text string, c chan<- string) {
	c <- text
}

func main() {

	// Despues de la coma se pone la cantidad maxima que recibe a la vez, si no se pone la cantidad es dinamica pero es menos optimo
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bay", c)

	fmt.Println(<-c)

}
