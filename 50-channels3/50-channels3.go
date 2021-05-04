package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	mensajes:= make(chan string, 4) //hasta 4 mensajes
	for  i:=0; i<5; i++ {
		go generarMensaje(mensajes, i)
	}
	imprimirMensaje(mensajes)
}

func generarMensaje(out chan<- string, num int) {
	cadena:="HELLO WORLD"
	for _,i:= range cadena{
		out<-string(i)+". Goroutine: "+strconv.Itoa(num)
	}
}

func imprimirMensaje(in <-chan string) {
	for i:=range in{
		fmt.Println(i)
		time.Sleep(100*time.Millisecond)
	}
}
