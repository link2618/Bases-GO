package main

import "fmt"

func main() {

	// Array
	var array [4]int
	array[0] = 1
	array[1] = 2

	fmt.Println(array)
	// Tamaño del array
	fmt.Println(len(array))
	// Capacidad maxima del array
	fmt.Println(cap(array))

	// Slice
	slice := []int{0, 1, 2, 3, 5, 6}
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// Metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:2])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
