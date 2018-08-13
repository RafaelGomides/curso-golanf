package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Fale cmg", 3)
	// fale("João", "Não dá", 1)

	go fale("Maria", "Fale cmg", 500)
	fale("João", "Não dá", 500)
	// fmt.Println("Teste")

	// time.Sleep(time.Second * 5)
	// fmt.Println("Fim")
}
