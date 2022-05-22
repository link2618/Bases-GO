package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB ram, %d GB disco y es %s", myPC.ram, myPC.disk, myPC.brand)
}

func main() {

	myPc := pc{ram: 16, disk: 20, brand: "msi"}
	fmt.Println(myPc)

}
