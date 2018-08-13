package main

import (
	"fmt"
	"time"
)

// Channel é a forma de comunicação entre as goroutines
// Channel é um tipo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // Enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c // recebendo dados do canal
	fmt.Println(a, b)
	fmt.Println(<-c)
}
