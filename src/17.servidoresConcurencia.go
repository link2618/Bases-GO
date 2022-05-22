package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()
	canal := make(chan string)

	servidores := []string{
		"http://google.com",
		"http://facebook.com",
		"http://Instagram.com",
	}

	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	tiempoTranscurrido := time.Since(inicio)

	fmt.Println("Tiempo de ejecucion: ", tiempoTranscurrido)

}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)

	if err != nil {
		fmt.Printf("Error server. ", err)
		canal <- servidor + " no se encuentra disponible"
	} else {
		fmt.Printf("El servidor %s esta disponible\n", servidor)
		canal <- servidor + " esta disponible"
	}
}
