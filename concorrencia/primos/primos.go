package main

import (
	"fmt"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				// time.Sleep(time.Millisecond * 10)
				break
			}
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 3000)
	go primos(cap(c), c)
	for primos := range c {
		fmt.Printf("%d ", primos)
	}
	fmt.Println()
}
