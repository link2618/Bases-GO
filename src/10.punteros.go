package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc, "PONG")
}

func (myPc *pc) duplicateRam() {
	myPc.ram = myPc.ram * 2
}

func main() {

	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	myPc := pc{ram: 16, disk: 20, brand: "msi"}
	fmt.Println(myPc)

	myPc.ping()
	myPc.duplicateRam()
	fmt.Println("Ultimo", myPc)

}
