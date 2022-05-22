package main

import "fmt"

// CarPublic Si empieza con mayuscula es publica
type CarPublic struct {
	Brand string
	Year  int
}

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
