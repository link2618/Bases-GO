package main

import "fmt"

func normalFunction(msg string) {
	fmt.Println(msg)
}

func tripeArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func dobleReturn(a int) (c, d int) {
	return a, d * 2
}

func main() {
	normalFunction("Hola mundo desde funcion")
	tripeArgument(1, 2, "hola")

	value := returnValue(12)
	fmt.Println(value)

	value1, value2 := dobleReturn(12)
	fmt.Println(value1, value2)

	value3, _ := dobleReturn(12)
	fmt.Println(value3)
}
