package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	// 01 - Mês
	// 02 - Dia
	// 2006 - Ano com 4 digitos
	// 03 - Hora AM/PM
	// 04 - Minuto
	// 05 - Segundos
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1> Hora Certa: %s</h1>", s)
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando")
	log.Fatal(http.ListenAndServe(":3333", nil))
}
