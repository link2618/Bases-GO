package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {

	myCar := car{brand: "VW", year: 2020}
	fmt.Println(myCar)

	// Instanciar clase vacia
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

}
