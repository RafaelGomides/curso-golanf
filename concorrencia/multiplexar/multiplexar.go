package main

import (
	"fmt"

	"github.com/RafaelGomides/html"
)

func encaminhar(ori <-chan string, des chan string) {
	for {
		des <- <-ori
	}
}

// Multiplexar - misturar (mensagens) num canal
func juntar(entr0, entr1 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entr0, c)
	go encaminhar(entr1, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.amazon.com", "https://www.youtube.com"),
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
