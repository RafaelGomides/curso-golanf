package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // Operação bloqueante
	fmt.Println("Só depois foi lido")
}

func main() {
	c := make(chan int) // canal sem buffer

	go rotina(c)

	fmt.Println(<-c) // Ação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c) // Deadlock
	fmt.Println("Fim")
}
