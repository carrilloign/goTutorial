package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go enviarPing(c)
	go recibirPing(c)

	fmt.Println("Presione enter para finalizar")
	var enter string
	fmt.Scanln(&enter)
}

func recibirPing(c chan string) {
	contador:=0
	for {
		fmt.Println(<-c,"",contador)
		contador++
		time.Sleep(1*time.Second)
	}
}

func enviarPing(c chan string) {
	for{
		c<-"Ping"
	}
}

