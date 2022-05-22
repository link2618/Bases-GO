package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()

	servidores := []string{
		"http://google.com",
		"http://facebook.com",
		"http://Instagram.com",
	}

	for _, servidor := range servidores {
		revisarServidor(servidor)
	}

	tiempoTranscurrido := time.Since(inicio)

	fmt.Println("Tiempo de ejecucion: ", tiempoTranscurrido)

}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)

	if err != nil {
		log.Fatal("Error server. ", err)
	} else {
		fmt.Printf("El servidor %s esta disponible\n", servidor)
	}
}
