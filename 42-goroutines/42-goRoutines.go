package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	conConcurrencia()
}

//utilizada para sincronizar las go routines
var wg sync.WaitGroup

func conConcurrencia() {
	wg.Add(2) //le indicamos al programa que va a trabajar con 2 goroutines
	fmt.Println("Iniciando goroutines")
	go imprimirCantidad("A")
	go imprimirCantidad("B")
	fmt.Println("Esperando que finalicen las go routines")
	wg.Wait() //en este punto main debe esperar por las goroutines
	fmt.Println("Programa finalizado")
}

func imprimirCantidad(s string) {
	for i := 0; i < 10; i++ {
		sleep:=rand.Int63n(1000)
		time.Sleep(time.Duration(sleep)* time.Millisecond)
		fmt.Printf("Vez %d imprimiendo con la etiqueta %s\n", i, s)
	}
	//Avisamos que la goroutine termino
	defer wg.Done()
}
