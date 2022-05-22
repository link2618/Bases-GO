package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["Hugo"] = 28
	m["Yenny"] = 25

	fmt.Println(m)

	// ok es para saber si la variable existe ya que en go si no existe retorna 0
	value, ok := m["Jose"]
	fmt.Println(value, ok)

}
